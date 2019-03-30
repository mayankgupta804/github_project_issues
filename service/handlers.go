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
	formValue := r.FormValue("repository_link")
	repositoryLink := strings.Split(formValue, "/")

	owner := repositoryLink[len(repositoryLink)-2]
	repoName := repositoryLink[len(repositoryLink)-1]

	repoInfo := owner + "," + repoName

	rc := GetRedisConnection()

	rc.Set(owner+repoName, []byte("incomplete"))
	// Publish the repoInfo to the github_service_queue for processing
	publisher([]byte(repoInfo))

	info := domain.RepositoryInfo{Owner: owner, Repository: repoName}

	// Set the header as accepted and render the loader template
	w.WriteHeader(http.StatusAccepted)
	tmpl := template.Must(template.ParseFiles("templates/loader.html"))
	tmpl.Execute(w, info)
}

// CheckStatus returns the status of completion of the background job
// which has the responsibility of fetching the issues related to a repository
func CheckStatus(w http.ResponseWriter, r *http.Request) {
	requestURI := r.RequestURI
	repositoryLink := strings.Split(requestURI, "/")

	owner := repositoryLink[len(repositoryLink)-2]
	repoName := repositoryLink[len(repositoryLink)-1]

	status := statusChecker(owner, repoName)
	mapStatus := make(map[string]string)
	if !status {
		mapStatus["status"] = "incomplete"
		json.NewEncoder(w).Encode(mapStatus)
		return
	}
	mapStatus["status"] = "complete"
	json.NewEncoder(w).Encode(mapStatus)
}

// DisplayGithubIssues renders a page with Github issues data for the relevant owner and repository
func DisplayGithubIssues(w http.ResponseWriter, r *http.Request) {
	requestURI := r.RequestURI
	repositoryLink := strings.Split(requestURI, "/")
	owner := repositoryLink[len(repositoryLink)-2]
	repoName := repositoryLink[len(repositoryLink)-1]

	issuesData := getIssuesData(owner, repoName)

	log.Printf("Received Data: %v", issuesData)

	if issuesData.Issues["Total Open Issues"] == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	tmpl.Execute(w, issuesData)
}
