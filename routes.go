package mux

import "net/http"

// Route contains the the details pertaining to an route
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes is a Route slice
type Routes []Route
