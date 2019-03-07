package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	owner := "smartystreets"
	repoName := "goconvey"
	totalOpenIssues := TotalOpenIssues(ctx, client, owner, repoName)
	issuesLast24Hours := OpenedInLast24Hours(ctx, client, owner, repoName)
	issuesLast7Days := OpenedInLast7Days(ctx, client, owner, repoName)
	issuesMoreThan7Days := totalOpenIssues - issuesLast7Days
	issuesMoreThan24HoursLessThan7Days := issuesLast7Days - issuesLast24Hours
	fmt.Printf("Total number of open issues: %d\n", totalOpenIssues)
	fmt.Printf("Total number of open issues in the last 24 hours: %d\n", issuesLast24Hours)
	fmt.Printf("Total number of issues opened more than 7 days ago: %d\n", issuesMoreThan7Days)
	fmt.Printf("Total number of issues opened more than 24 hours ago but less than 7 days ago: %d\n", issuesMoreThan24HoursLessThan7Days)
}

// TotalOpenIssues returns total open issues of a repository
func TotalOpenIssues(ctx context.Context, client *github.Client, owner string, repoName string) int {
	opts := &github.IssueListByRepoOptions{State: "open"} // Default option for State is 'open'
	// opts := &github.IssueListByRepoOptions{State: 'open'} would work the same way as the above
	issues, resp, err := client.Issues.ListByRepo(ctx, owner, repoName, opts) // Returns a list of all the open issues
	if err != nil {
		log.Fatalf("Error Encountered: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Couldn't get the list of open issues. Please try again after some time.")
	}
	return len(issues)
}

// OpenedInLast24Hours returns total issues opened in the last 24 hours
func OpenedInLast24Hours(ctx context.Context, client *github.Client, owner string, repoName string) int {
	now := time.Now()
	after := now.AddDate(0, 0, -1)
	opts := &github.IssueListByRepoOptions{Since: after}
	issues, resp, err := client.Issues.ListByRepo(ctx, owner, repoName, opts)
	if err != nil {
		log.Fatalf("Error Encountered: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Couldn't get the list of open issues. Please try again after some time.")
	}
	return len(issues)
}

// OpenedInLast7Days returns total issues opened more than 7 days ago
func OpenedInLast7Days(ctx context.Context, client *github.Client, owner string, repoName string) int {
	now := time.Now()
	after := now.AddDate(0, 0, -7)
	opts := &github.IssueListByRepoOptions{Since: after}
	issues, resp, err := client.Issues.ListByRepo(ctx, owner, repoName, opts)
	if err != nil {
		log.Fatalf("Error Encountered: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Couldn't get the list of open issues. Please try again after some time.")
	}
	return len(issues)
}
