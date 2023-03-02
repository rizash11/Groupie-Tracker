package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	app := application{
		info_log:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		error_log: log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	app.Unmarshal()

	my_server := &http.Server{
		Addr:    ":4001",
		Handler: app.router(),
	}

	app.info_log.Println("The server is starting at http://localhost:4001/")
	app.error_log.Fatalln(my_server.ListenAndServe())
}

func (app *application) Unmarshal() {
	// ARTISTS
	data, err := os.ReadFile("./backend/api/artists.txt")
	if err != nil {
		app.error_log.Fatalln(err)
	}

	err = json.Unmarshal(data, &app.artists)
	if err != nil {
		app.error_log.Fatalln(err)
	}

	// DATES
	data, err = os.ReadFile("./backend/api/dates.txt")
	if err != nil {
		app.error_log.Fatalln(err)
	}

	err = json.Unmarshal(data, &app.dates)
	if err != nil {
		app.error_log.Fatalln(err)
	}

	// LOCATIONS
	data, err = os.ReadFile("./backend/api/locations.txt")
	if err != nil {
		app.error_log.Fatalln(err)
	}

	err = json.Unmarshal(data, &app.locations)
	if err != nil {
		app.error_log.Fatalln(err)
	}
}
