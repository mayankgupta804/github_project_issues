package main

import (
	"log"
	"net/http"

	"github.com/radius_agents_assignment/github_project_issues/router"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
