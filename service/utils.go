package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/radius_agents_assignment/github_project_issues/domain"
	"github.com/radius_agents_assignment/github_project_issues/queue"
	"github.com/radius_agents_assignment/github_project_issues/redisclient"
)

// publisher publishes a job to the background worker to process
func publisher(repoinfo []byte) {
	if err := queue.Publish("github_service_queue", repoinfo); err != nil {
		panic(err)
	}
}

// statusChecker checks the status of a given task in redis
func statusChecker(owner string, repository string) bool {
	rc := GetRedisConnection()
	data, err := rc.Get(owner + repository)
	if err != nil {
		log.Printf("Error Encountered: %v", err)
		return false
	}
	issuesData := domain.IssuesData{}
	if err = json.Unmarshal([]byte(data), &issuesData); err != nil {
		return false
	}
	return true
}

// getIssuesData returns the processed data from redis for a particular owner and repository
func getIssuesData(owner string, repository string) *domain.IssuesData {
	rc := GetRedisConnection()
	data, err := rc.Get(owner + repository)
	issuesData := domain.IssuesData{}
	if err = json.Unmarshal([]byte(data), &issuesData); err != nil {
		issuesData.Issues["Total Open Issues"] = 0
		return &issuesData
	}
	return &issuesData
}

// GetRedisConnection returns a redis client with a connection to the redis pool
func GetRedisConnection() *redisclient.RedisClient {
	redisHostURL := os.Getenv("REDIS_URL")
	if redisHostURL == "" {
		log.Fatal("$REDIS_URL must be set")
	}
	redisClient := redisclient.RedisClient{}
	redisClient.Init(redisHostURL)
	return &redisClient
}
