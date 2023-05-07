package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (app *application) error_checker(err error) {
	if err != nil {
		app.error_log.Fatalln(err)
	}
}

func (app *application) Unmarshal() {
	// ARTISTS
	data, err := os.ReadFile("./backend/api/artists.txt")
	app.error_checker(err)

	err = json.Unmarshal(data, &app.td.Artists)
	app.error_checker(err)

	for i, artist := range app.td.Artists {
		first_album_date := strings.Split(artist.FirstAlbum, "-")
		app.td.Artists[i].FirstAlbum = first_album_date[2] + "-" + first_album_date[1] + "-" + first_album_date[0]
	}

	// DATES
	data, err = os.ReadFile("./backend/api/dates.txt")
	app.error_checker(err)

	err = json.Unmarshal(data, &app.td.Dates)
	app.error_checker(err)

	// LOCATIONS
	data, err = os.ReadFile("./backend/api/locations.txt")
	app.error_checker(err)

	err = json.Unmarshal(data, &app.td.Locations)
	app.error_checker(err)

	for i := range app.td.Locations.Index {
		var location_split []string
		for j, location := range app.td.Locations.Index[i].Locations {
			location = strings.Replace(location, "_", " ", -1)
			location_split = strings.Split(location, "-")
			caser := cases.Title(language.English)
			location_split[0] = caser.String(location_split[0])
			location_split[1] = caser.String(location_split[1])

			app.td.Locations.Index[i].Locations[j] = location_split[0] + ", " + location_split[1]

			app.locations_filter2[location_split[1]] = map[string][]*artist{}
			app.locations_filter2[location_split[1]][location_split[0]] = append(app.locations_filter2[location_split[1]][location_split[0]], &app.td.Artists[i])
		}

		app.locations_filter1[location_split[1]] = append(app.locations_filter1[location_split[1]], &app.td.Artists[i])
	}

	// RELATIONS
	data, err = os.ReadFile("./backend/api/relation.txt")
	app.error_checker(err)

	err = json.Unmarshal(data, &app.td.Relations)
	app.error_checker(err)
}

func (app *application) AddLinks() {
	for i, artist := range app.td.Artists {
		link := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(artist.Id)
		app.td.Artists[i].Artists_link = template.URL(link)
	}
}

func (app *application) NotFound(w http.ResponseWriter, r *http.Request) {
	// app.error_log.Println(http.StatusText(http.StatusNotFound))
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.error_log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (td *template_data) Add(x, y int) int {
	return x + y
}

func (nfs *neutered_fs) Open(dir string) (http.File, error) {
	f, err := nfs.fs.Open(dir)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(dir, "index.html")
		if _, err = nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
