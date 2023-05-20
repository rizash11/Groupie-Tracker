package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) filter(r *http.Request) {
	// Cleaning up after previous filters
	for i := range app.td.Artists {
		app.td.Artists[i].Show = false
	}
	app.td.Filter_errors = filter_error{
		Cdate_err:      nil,
		Firstalbum_err: nil,
		Members_err:    nil,
		Locations_err:  nil,
	}

	r.ParseForm()

	app.td.Filter_errors.Cdate_err = app.filter_by_cdate(r)
	app.td.Filter_errors.Firstalbum_err = app.filter_by_firstalbum(r)
	app.td.Filter_errors.Members_err = app.filter_by_members(r)
	app.td.Filter_errors.Locations_err = app.filter_by_locations(r)

	for i := range app.td.Artists {
		f1 := app.td.Artists[i].filters.show_by_cdate
		f2 := app.td.Artists[i].filters.show_by_first_album
		f3 := app.td.Artists[i].filters.show_by_members
		f4 := app.td.Artists[i].filters.show_by_location

		if f1 && f2 && f3 && f4 {
			app.td.Artists[i].Show = true
		}
	}
}

func (app *application) filter_by_cdate(r *http.Request) error {
	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_cdate = true
	}

	var c_date_filter_1 int
	var err error
	input_from := r.FormValue("c-date-filter-1")
	if input_from != "" {
		c_date_filter_1, err = strconv.Atoi(input_from)
		if err != nil {
			return err
		} else if c_date_filter_1 > 3000 {
			return errors.New("dude wtf?")
		}
	}

	var c_date_filter_2 int
	input_to := r.FormValue("c-date-filter-2")
	if input_to != "" {
		c_date_filter_2, err = strconv.Atoi(input_to)
		if err != nil {
			return err
		}
	} else {
		c_date_filter_2 = 3000
	}

	if c_date_filter_1 > c_date_filter_2 {
		return errors.New("date \"from\" must not be bigger than date \"to\"")
	}

	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_cdate = false
	}

	for i, artist := range app.td.Artists {
		if c_date_filter_1 <= artist.CreationDate && artist.CreationDate <= c_date_filter_2 {
			app.td.Artists[i].filters.show_by_cdate = true
		}
	}

	return nil
}

func (app *application) filter_by_firstalbum(r *http.Request) error {
	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_first_album = true
	}

	first_album_1 := r.FormValue("first-album-1")
	if first_album_1 == "" {
		first_album_1 = "0000-00-00"
	} else if first_album_1 > "3000-01-01" {
		return errors.New("dude what in the hell... ")
	}

	first_album_2 := r.FormValue("first-album-2")
	if first_album_2 == "" {
		first_album_2 = "3000-01-01"
	}

	if first_album_1 > first_album_2 {
		return errors.New("first album: date from must not be bigger than date to")
	}

	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_first_album = false
	}

	for i, artist := range app.td.Artists {
		if first_album_1 <= artist.FirstAlbum && artist.FirstAlbum <= first_album_2 {
			app.td.Artists[i].filters.show_by_first_album = true
		}
	}

	return nil
}

func (app *application) filter_by_members(r *http.Request) error {
	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_members = true
	}

	var members int
	var err error
	input := r.FormValue("members-filter")
	if input != "" {
		members, err = strconv.Atoi(input)
		if err != nil {
			return err
		} else if members > 3000 {
			return errors.New("dude wtf?")
		}
	} else {
		return nil
	}

	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_members = false
	}

	for i, artist := range app.td.Artists {
		if members == len(artist.Members) {
			app.td.Artists[i].filters.show_by_members = true
		}
	}

	return nil
}

func (app *application) filter_by_locations(r *http.Request) error {
	for i := range app.td.Artists {
		app.td.Artists[i].filters.show_by_location = true
	}

	countries_cities := r.Form["city"]
	if len(countries_cities) != 0 {
		for i := range app.td.Artists {
			app.td.Artists[i].filters.show_by_location = false
		}
	}

	for _, country_city := range countries_cities {
		split := strings.Split(country_city, "-")
		for _, link := range app.locations_filter2[split[0]][split[1]] {
			link.filters.show_by_location = true
		}
	}

	return nil
}
