package service

import (
	"encoding/json"
	"log"

	"github.com/radius_agents_assignment/github_project_issues/domain"
	"github.com/radius_agents_assignment/github_project_issues/queue"
)

func publisher(repoinfo []byte) {
	if err := queue.Publish("github_service_queue", repoinfo); err != nil {
		panic(err)
	}
}

func statusChecker() bool {
	msgs, close, err := queue.Subscribe("github_service_consume_queue")

	if err != nil {
		panic(err)
	}
	defer close()

	data := domain.IssuesData{}

	for d := range msgs {

		err = json.Unmarshal(d.Body, &data)

		if err != nil {
			log.Printf("Error converting JSON object: %v", err)
			return false
		}
		return true
	}
	return false
}

func subscriber() *domain.IssuesData {
	// Subscribe to the queue where the worker publishes the result
	msgs, close, err := queue.Subscribe("github_service_consume_queue")
	if err != nil {
		panic(err)
	}
	// Close the channel after getting the results
	defer close()

	data := domain.IssuesData{}

	for d := range msgs {

		err = json.Unmarshal(d.Body, &data)

		if err != nil {
			data.Issues = map[string]int{"Total Open Issues": 0}
			return &data
		}

		// Acknowledge the message so that it is cleared from the queue
		d.Ack(true)

		return &data
	}
	// If not messages are present, send a partially empty struct with Issues map
	data.Issues = map[string]int{"Total Open Issues": 0}
	return &data
}
