package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radius_agents_assignment/github_project_issues/routes"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.Routes {
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
