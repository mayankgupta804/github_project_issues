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

	// Create token source to pass the access token to the client for authentication
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)

	// Create a new oauth client using the token source and backgrdound context
	tc := oauth2.NewClient(ctx, ts)

	// Create a new github client to communicate with the Github API
	client = github.NewClient(tc)
	return client, ctx
}
