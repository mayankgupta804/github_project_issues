package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Github Issues Viewer</h1>")
}

func GithubIssues(w http.ResponseWriter, r *http.Request) {
	var owner = mux.Vars(r)["owner"]
	var repoName = mux.Vars(r)["repoName"]
	issuesMap, err := GetGithubIssues(owner, repoName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(issuesMap)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
