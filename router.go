package mux

import (
	"strings"

	"github.com/gorilla/mux"
)

// NewRouter constructs a new mux router with a slice of user defined routes
func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		methods := strings.Split(route.Method, ",")
		router.
			Methods(methods...).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
