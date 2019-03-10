package domain

// IssuesData is structure for storing data about the issues of a Github repository and
// It is being used for the purpose of displaying the data in layout.html
type IssuesData struct {
	Owner      string
	Repository string
	Issues     map[string]int
}
