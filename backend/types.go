package main

import (
	"log"
)

type artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type application struct {
	info_log  log.Logger
	error_log log.Logger
	artists   []artist
}
