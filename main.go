package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

const (
	today        = 0
	oneDayAgo    = -1
	sevenDaysAgo = -7
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		jsonStr := getGithubIssues()

		c.String(http.StatusOK, string(jsonStr))
	})
	router.Run(":8080")
}

func getGithubIssues() []byte {
	client, ctx := authenticateClient()
	owner := "smartystreets"
	repoName := "goconvey"
	totalOpenIssues := repoIssuesCounter(ctx, client, owner, repoName, today)
	issuesLast24Hours := repoIssuesCounter(ctx, client, owner, repoName, oneDayAgo)
	issuesLast7Days := repoIssuesCounter(ctx, client, owner, repoName, sevenDaysAgo)
	issuesMoreThan7Days := totalOpenIssues - issuesLast7Days
	issuesMoreThan24HoursLessThan7Days := issuesLast7Days - issuesLast24Hours
	fmt.Printf("Total number of open issues: %d\n", totalOpenIssues)
	fmt.Printf("Total number of open issues in the last 24 hours: %d\n", issuesLast24Hours)
	fmt.Printf("Total number of issues opened more than 7 days ago: %d\n", issuesMoreThan7Days)
	fmt.Printf("Total number of issues opened more than 24 hours ago but less than 7 days ago: %d\n", issuesMoreThan24HoursLessThan7Days)

	issuesMap := make(map[string]int)
	issuesMap["totalOpenIssues"] = totalOpenIssues
	issuesMap["issuesLast24Hours"] = issuesLast24Hours
	issuesMap["issuesMoreThan7DaysAgo"] = issuesMoreThan7Days
	issuesMap["issuesMoreThan24HoursLessThan7Days"] = issuesMoreThan24HoursLessThan7Days

	data, err := json.Marshal(issuesMap)
	if err != nil {
		log.Fatalf("Unable to convert map object to a JSON string: %v", err)
	}
	return data
}

func authenticateClient() (*github.Client, context.Context) {
	username := "mayankgupta804"
	password := "satishgupta52"
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	ctx := context.Background()
	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		log.Printf("\nerror: %v\n", err)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
	return client, ctx
}

func repoIssuesCounter(ctx context.Context, client *github.Client, owner string, repoName string, daysAgo int) int {
	var options = &github.IssueListByRepoOptions{State: "open"}
	if daysAgo < 0 {
		now := time.Now()
		sinceTime := now.AddDate(0, 0, daysAgo)
		options.Since = sinceTime
	}
	var issuesCount, nextPage = 0, 0
	var issues []*github.Issue
	var resp *github.Response
	var err error
	for {
		issues, resp, err = client.Issues.ListByRepo(ctx, owner, repoName, options)
		nextPage = resp.NextPage
		issuesCount += len(issues)

		if nextPage == 0 {
			return issuesCount
		}

		options.Page = nextPage

		if err != nil {
			log.Fatalf("Error Encountered: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Println("Couldn't get the list of open issues. Please try again after some time.")
		}
	}
}
