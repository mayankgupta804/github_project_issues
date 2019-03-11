package service

import client "github.com/radius_agents_assignment/github_project_issues/githubclient"

// GetGithubIssues returns a map object and error (if any) with information related to total open issues,
// issues opened in last 24 hours, issues opened more than 7 days ago and
// issues opened more than 24 hours ago but less than 7 days ago for a given repository
func GetGithubIssues(owner string, repoName string) map[string]int {
	client, ctx := client.GithubClientContext()

	totalOpenIssues, _ := repoIssuesCounter(ctx, client, owner, repoName, today)
	issuesLast24Hours, _ := repoIssuesCounter(ctx, client, owner, repoName, oneDayAgo)
	issuesLast7Days, _ := repoIssuesCounter(ctx, client, owner, repoName, sevenDaysAgo)

	issuesMoreThan7Days := totalOpenIssues - issuesLast7Days
	issuesMoreThan24HoursLessThan7Days := issuesLast7Days - issuesLast24Hours

	issuesMap := make(map[string]int)
	issuesMap["Total Open Issues"] = totalOpenIssues
	issuesMap["Issues Opened in Last 24 Hours"] = issuesLast24Hours
	issuesMap["Issues Opened More Than 7 Days Ago"] = issuesMoreThan7Days
	issuesMap["Issues Opened More Than 24Hours Ago But Less Than 7 Days Ago"] = issuesMoreThan24HoursLessThan7Days

	return issuesMap
}
