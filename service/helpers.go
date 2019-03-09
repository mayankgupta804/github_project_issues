package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
)

// repoIssuesCounter counts issues in a repository
func repoIssuesCounter(ctx context.Context, client *github.Client, owner string, repoName string, daysAgo int) int {
	options := getIssueListByRepoOptions(daysAgo) // get options for querying data from github issues API

	var issuesCount, nextPage = 0, 0
	var issues []*github.Issue // slice object which holds a list of all github issues object
	var resp *github.Response  // wrapper over a normal http.Response object
	var err error

	for {
		issues, resp, err = client.Issues.ListByRepo(ctx, owner, repoName, options)
		nextPage = resp.NextPage // store the index of next page in case of paginated results
		issuesCount += len(issues)

		if nextPage == 0 {
			return issuesCount // in case there are no further paginated results, return the final count
		}

		if err != nil {
			log.Fatalf("Error Encountered: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Println("Couldn't get the list of open issues. Please try again after some time.")
		}

		options.Page = nextPage // set the index of the next page of results in options
	}
}

// getIssuesListByRepoOptions provides options according to which the issues will be fetched
// from a given repository
func getIssueListByRepoOptions(daysAgo int) *github.IssueListByRepoOptions {
	var options = &github.IssueListByRepoOptions{State: "open"} // only fetch issues that are "open"
	if daysAgo < 0 {
		now := time.Now()
		sinceTime := now.AddDate(0, 0, daysAgo)
		options.Since = sinceTime
	}
	return options
}
