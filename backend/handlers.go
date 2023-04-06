package main

import (
	"html/template"
	"net/http"
)

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)

	file_server := http.FileServer(&neutered_fs{http.Dir("./frontend/")})
	mux.Handle("/frontend/", http.StripPrefix("/frontend", file_server))

	return mux
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./frontend/html/home.page.html", "./frontend/html/base.layout.html", "./frontend/html/header.partial.html", "./frontend/html/footer.partial.html")
	if err != nil {
		app.ServerError(w, r, err)
		return
	}

	err = tmpl.Execute(w, app.td)
	if err != nil {
		app.ServerError(w, r, err)
		return
	}
}
