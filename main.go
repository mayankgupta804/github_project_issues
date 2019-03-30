package main

import (
	"log"
	"os"

	"github.com/radius_agents_assignment/github_project_issues/queue"
	"github.com/radius_agents_assignment/github_project_issues/service"
	"github.com/radius_agents_assignment/github_project_issues/worker"
)

var appName = "github_issues_service"

func main() {
	rabbitMQHostAddress := os.Getenv("CLOUDAMQP_URL")
	if rabbitMQHostAddress == "" {
		log.Fatal("$RABBITMQ_HOST_ADDRESS must be set")
	}
	queue.Init(rabbitMQHostAddress)

	if os.Args[1] == "worker" {
		log.Println("Starting worker...")
		worker.StartWorker()
	} else if os.Args[1] == "server" {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("$PORT must be set")
		}
		log.Printf("Starting %s...\n", appName)
		service.StartWebServer(port)
	}
}
