package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/radius_agents_assignment/github_project_issues/queue"
	"github.com/radius_agents_assignment/github_project_issues/service"
)

// StartWorker starts the Go worker process for making API calls on behalf of the web application
func StartWorker() {
	// Get the channel which we subscribe to
	msgs, close, err := queue.Subscribe("github_service_queue")
	if err != nil {
		panic(err)
	}
	// Close the channel when the worker is stopped
	defer close()

	stop := make(chan bool)
	var repoInfoDataBytes []byte

	rc := service.GetRedisConnection()

	go func() {
		// Receive messages from the channel forever
		for d := range msgs {
			// When a message a received, pass it is an argument to GithubIssuesFetcher to crunch the issues data
			repoInfo := service.GithubIssuesFetcher(d.Body)

			log.Println(time.Now().Format("01-02-2006 15:04:05"), "::", repoInfo)

			repoInfoDataBytes, _ = json.Marshal(repoInfo)

			// Push the data to redis after job has been completed
			if err := rc.Set(repoInfo.Owner+repoInfo.Repository, repoInfoDataBytes); err != nil {
				log.Fatalf("Error Encountered while storing data in Redis: %v", err)
			}

			// Acknowledge the message so that it is cleared from the queue
			d.Ack(true)
		}
	}()
	log.Println("To exit press CTRL+C")
	<-stop
}
