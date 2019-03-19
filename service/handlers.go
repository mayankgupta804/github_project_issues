package service

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/radius_agents_assignment/github_project_issues/domain"
)

// Index renders a page with an HTML form for entering the repository name and owner/organisation name
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

// GetGithubIssues renders the page with the relevant information about the issues related to a
// specific repository
func GetGithubIssues(w http.ResponseWriter, r *http.Request) {
	repositoryLink := strings.Split(r.FormValue("repository_link"), "/")
	owner := repositoryLink[len(repositoryLink)-2]
	repoName := repositoryLink[len(repositoryLink)-1]
	repoInfo := owner + "," + repoName

	// Publish the repoInfo to the github_service_queue for processing
	publisher([]byte(repoInfo))

	w.WriteHeader(http.StatusAccepted)
	tmpl := template.Must(template.ParseFiles("templates/loader.html"))
	tmpl.Execute(w, nil)
}

// CheckStatus returns the status of the background job of fetching the issues related to a repository
func CheckStatus(w http.ResponseWriter, r *http.Request) {
	status := statusChecker()
	var mapStatus map[string]string
	if !status {
		mapStatus = map[string]string{"status": "incomplete"}
		json.NewEncoder(w).Encode(mapStatus)
		return
	}
	mapStatus["status"] = "complete"
	json.NewEncoder(w).Encode(mapStatus)
}

// DisplayGithubIssues displays a page with the data about the issues
func DisplayGithubIssues(w http.ResponseWriter, r *http.Request) {
	data := make(chan map[string]int)
	go subscriber(data)
	issuesMap := <-data
	log.Printf("Received Data: %v", issuesMap)

	if issuesMap["Total Open Issues"] == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	issuesData := domain.IssuesData{Owner: "owner", Repository: "repoName", Issues: issuesMap}
	tmpl.Execute(w, issuesData)
}
