package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/radius_agents_assignment/github_project_issues/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func GithubIssues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/code; charset=UTF-8")

	jsonStr := service.GetGithubIssues()
	data := make(map[string]int)
	json.Unmarshal(jsonStr, &data)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
