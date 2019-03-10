package main

import (
	"fmt"
	"os"

	"github.com/radius_agents_assignment/github_project_issues/service"
)

var appName = "github_issues_service"

func main() {
	fmt.Printf("Starting %s\n", appName)
	port := os.Getenv("PORT")
	service.StartWebServer(port)
}
