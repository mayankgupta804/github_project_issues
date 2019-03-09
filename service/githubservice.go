package service

import (
	"encoding/json"
	"log"
)

// GetGithubIssues returns a JSON object with information related to total open issues,
// issues opened in last 24 hours, issues opened more than 7 days ago and
// issues opened more than 24 hours ago but less than 7 days ago for a given repository
func GetGithubIssues() []byte {
	// client, ctx := client.GithubClientContext()
	// owner := "smartystreets"
	// repoName := "goconvey"

	// totalOpenIssues := repoIssuesCounter(ctx, client, owner, repoName, today)
	// issuesLast24Hours := repoIssuesCounter(ctx, client, owner, repoName, oneDayAgo)
	// issuesLast7Days := repoIssuesCounter(ctx, client, owner, repoName, sevenDaysAgo)
	// issuesMoreThan7Days := totalOpenIssues - issuesLast7Days
	// issuesMoreThan24HoursLessThan7Days := issuesLast7Days - issuesLast24Hours

	issuesMap := make(map[string]int)
	issuesMap["totalOpenIssues"] = 5
	issuesMap["issuesLast24Hours"] = 6
	issuesMap["issuesMoreThan7DaysAgo"] = 10
	issuesMap["issuesMoreThan24HoursLessThan7Days"] = 12

	data, err := json.Marshal(issuesMap)
	if err != nil {
		log.Fatalf("Unable to convert map object to a JSON string: %v", err)
	}
	return data
}
