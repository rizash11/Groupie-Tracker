package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
