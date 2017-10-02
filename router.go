package mux

import (
	"net/http"
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

// Handle is a middleware function that executes operations on a request before passing it to the main handler
func Handle(handler http.Handler, operations ...Operation) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, operation := range operations {
			if !operation(w, r) {
				return
			}
		}

		handler.ServeHTTP(w, r)
	})
}
