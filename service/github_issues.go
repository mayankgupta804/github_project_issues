package service

import client "github.com/radius_agents_assignment/github_project_issues/githubclient"

// GetGithubIssues returns a JSON object with information related to total open issues,
// issues opened in last 24 hours, issues opened more than 7 days ago and
// issues opened more than 24 hours ago but less than 7 days ago for a given repository
func GetGithubIssues(owner string, repoName string) (map[string]int, error) {
	client, ctx := client.GithubClientContext()

	totalOpenIssues, err := repoIssuesCounter(ctx, client, owner, repoName, today)
	if err != nil {
		return make(map[string]int), err
	}

	issuesLast24Hours, err := repoIssuesCounter(ctx, client, owner, repoName, oneDayAgo)
	if err != nil {
		return make(map[string]int), err
	}

	issuesLast7Days, err := repoIssuesCounter(ctx, client, owner, repoName, sevenDaysAgo)
	if err != nil {
		return make(map[string]int), err
	}

	issuesMoreThan7Days := totalOpenIssues - issuesLast7Days
	issuesMoreThan24HoursLessThan7Days := issuesLast7Days - issuesLast24Hours

	issuesMap := make(map[string]int)
	issuesMap["totalOpenIssues"] = totalOpenIssues
	issuesMap["issuesLast24Hours"] = issuesLast24Hours
	issuesMap["issuesMoreThan7DaysAgo"] = issuesMoreThan7Days
	issuesMap["issuesMoreThan24HoursLessThan7Days"] = issuesMoreThan24HoursLessThan7Days

	return issuesMap, nil
}
