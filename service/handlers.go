package service

import (
	"html/template"
	"net/http"
)

// Index renders a page with an HTML form for entering the repository name and owner/organisation name
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

// GithubIssues renders the page with the relevant information about the issues related to a
// specific repository
func GithubIssues(w http.ResponseWriter, r *http.Request) {
	owner := r.FormValue("owner")
	repoName := r.FormValue("repoName")

	issuesMap, err := GetGithubIssues(owner, repoName)

	if err != nil || issuesMap["Total Open Issues"] == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	issuesData := IssuesData{Owner: owner, Repository: repoName, Issues: issuesMap}
	tmpl.Execute(w, issuesData)
}

// IssuesData is structure for storing data about the issues of a Github repository and
// It is being used for the purpose of displaying the data in layout.html
type IssuesData struct {
	Owner      string
	Repository string
	Issues     map[string]int
}
