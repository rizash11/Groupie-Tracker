package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	app := application{
		info_log:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		error_log: log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		td:        &template_data{},
	}

	app.Unmarshal()
	app.AddLinks()

	my_server := &http.Server{
		Addr:    ":4001",
		Handler: app.router(),
	}

	app.info_log.Println("The server is starting at http://localhost:4001/")
	app.error_log.Fatalln(my_server.ListenAndServe())
}
