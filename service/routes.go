package service

import (
	"net/http"
)

// Route defines a structure for defining API routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes stores a slice of routes which is used when
// when an instance of NewRouter is called to initialize the routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},
	Route{
		"Issues",
		http.MethodGet,
		"/issues",
		GetGithubIssues,
	},
	Route{
		"Status",
		http.MethodGet,
		"/status",
		CheckStatus,
	},
	Route{
		"IssuesData",
		http.MethodGet,
		"/issues/data",
		DisplayGithubIssues,
	},
}
