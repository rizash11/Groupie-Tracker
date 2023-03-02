package main

import (
	"log"
)

type application struct {
	info_log  *log.Logger
	error_log *log.Logger
	artists   []artist
	dates     struct_dates
	locations struct_locations
}

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

type struct_dates struct {
	Index []date
}

type date struct {
	Id    int
	Dates []string
}

type struct_locations struct {
	Index []location
}

type location struct {
	Id        int
	Locations []string
	Dates     string
}

type struct_relations struct {
	Index []relation
}

type relation struct {
	Id        int
	Locations []string
	Dates     string
}
