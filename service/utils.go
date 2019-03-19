package service

import (
	"encoding/json"
	"log"

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

	if len(msgs) > 0 {
		return true
	}
	return false
}

func subscriber(recvData chan<- map[string]int) {
	msgs, close, err := queue.Subscribe("github_service_consume_queue")
	if err != nil {
		panic(err)
	}
	// Close the channel when the worker is stopped
	defer close()
	var data map[string]int

	for d := range msgs {

		err = json.Unmarshal(d.Body, &data)

		if err != nil {
			recvData <- map[string]int{"Total Open Issues": 0}
			log.Fatalf("Error encountered: %s", err)
		}

		// Acknowledge the message so that it is cleared from the queue
		d.Ack(true)

		recvData <- data
	}
	recvData <- map[string]int{"Total Open Issues": 0}
}
