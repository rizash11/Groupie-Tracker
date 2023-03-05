package main

import (
	"html/template"
	"net/http"
)

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)

	return mux
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./frontend/html/base.layout.html", "./frontend/html/home.page.html")
	if err != nil {
		app.ServerError(w, r, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		app.ServerError(w, r, err)
		return
	}
}
