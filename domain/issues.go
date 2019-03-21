package domain

// IssuesData is a structure for storing data of a Github repository and the relevant issues
type IssuesData struct {
	Owner      string
	Repository string
	Issues     map[string]int
}
