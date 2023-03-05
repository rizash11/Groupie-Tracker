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
	relations struct_relations
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
	Index []index_date
}

type index_date struct {
	Id    int
	Dates []string
}

type struct_locations struct {
	Index []index_location
}

type index_location struct {
	Id        int
	Locations []string
	Dates     string
}

type struct_relations struct {
	Index []index_relation
}

type index_relation struct {
	Id             int
	DatesLocations map[string][]string
}
