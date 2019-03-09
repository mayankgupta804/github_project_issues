package routes

import (
	"net/http"

	"github.com/radius_agents_assignment/github_project_issues/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []Route

var Routes = routes{
	Route{
		"Index",
		http.MethodGet,
		"/",
		handlers.Index,
	},
	Route{
		"Issues",
		http.MethodGet,
		"/issues",
		handlers.GithubIssues,
	},
}
