package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/radius_agents_assignment/github_project_issues/domain"
	"github.com/radius_agents_assignment/github_project_issues/queue"
	"github.com/radius_agents_assignment/github_project_issues/redisclient"
)

func publisher(repoinfo []byte) {
	if err := queue.Publish("github_service_queue", repoinfo); err != nil {
		panic(err)
	}
}

func statusChecker(owner string, repository string) bool {
	rc := GetRedisConnection()
	data, err := rc.Get(owner + repository)
	if err != nil {
		return false
	}
	issuesData := domain.IssuesData{}
	if err = json.Unmarshal([]byte(data), &issuesData); err != nil {
		return false
	}
	return true
}

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

func GetRedisConnection() *redisclient.RedisClient {
	redisHostURL := os.Getenv("REDIS_URL")
	if redisHostURL == "" {
		log.Fatal("$REDIS_URL must be set")
	}
	redisClient := redisclient.RedisClient{}
	redisClient.Init(redisHostURL)
	return &redisClient
}
