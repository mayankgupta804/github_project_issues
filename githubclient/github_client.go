package client

import (
	"context"

	"github.com/google/go-github/github"
)

// GithubClientContext provides a client and background context to use the Github API
func GithubClientContext() (*github.Client, context.Context) {
	client := github.NewClient(nil)
	ctx := context.Background()
	return client, ctx
}
