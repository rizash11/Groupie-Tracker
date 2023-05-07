package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	info_log          *log.Logger
	error_log         *log.Logger
	td                *template_data
	locations_filter1 map[string][]*artist
	locations_filter2 map[string]map[string][]*artist
}

type artist struct {
	Id           int
	Image        template.URL
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    template.URL
	ConcertDates template.URL
	Relations    template.URL
	Artists_link template.URL
	Show         bool
	filters      struct_filters
}

type template_data struct {
	Artists       []artist
	Dates         struct_dates
	Locations     struct_locations
	Relations     struct_relations
	Filter_errors filter_error
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

type neutered_fs struct {
	fs http.FileSystem
}

type struct_filters struct {
	show_by_cdate       bool
	show_by_first_album bool
	show_by_members     bool
	show_by_location    bool
}

type filter_error struct {
	Cdate_err      error
	Firstalbum_err error
	Members_err    error
	Locations_err  error
}
