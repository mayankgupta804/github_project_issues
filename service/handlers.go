package service

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/radius_agents_assignment/github_project_issues/domain"
)

// Index renders a page with an HTML form for entering the repository name and owner/organisation name
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

// GithubIssues renders the page with the relevant information about the issues related to a
// specific repository
func GithubIssues(w http.ResponseWriter, r *http.Request) {
	repositoryLink := strings.Split(r.FormValue("repository_link"), "/")
	owner := repositoryLink[len(repositoryLink)-2]
	repoName := repositoryLink[len(repositoryLink)-1]
	issuesMap, err := GetGithubIssues(owner, repoName)

	if err != nil || issuesMap["Total Open Issues"] == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	issuesData := domain.IssuesData{Owner: owner, Repository: repoName, Issues: issuesMap}
	tmpl.Execute(w, issuesData)
}
