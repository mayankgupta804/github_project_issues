package client

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubClientContext provides a client and background context to use the Github API
func GithubClientContext() (*github.Client, context.Context) {
	ctx := context.Background()
	var client *github.Client

	// if os.Getenv("GITHUB_ACCESS_TOKEN") == "" { // If personal access token is not present, use basic client
	// 	log.Println("Github personal access token is not set. Please set it.")
	// 	client = github.NewClient(nil)
	// 	return client, ctx
	// }

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
	return client, ctx
}
