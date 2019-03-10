package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns a router object which sets up the routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(handler)
	}
	return router
}
