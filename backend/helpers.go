package main

import (
	"encoding/json"
	"net/http"
	"os"
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
