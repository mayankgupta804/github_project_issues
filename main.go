package main

import (
	"fmt"
	"log"
	"os"

	"github.com/radius_agents_assignment/github_project_issues/service"
)

var appName = "github_issues_service"

func main() {
	fmt.Printf("Starting %s\n", appName)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	service.StartWebServer(port)
}
