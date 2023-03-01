package main

import "net/http"

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/home/", app.home)

	return mux
}

func (app *application) home(w http.ResponseWriter, req *http.Request) {

}
