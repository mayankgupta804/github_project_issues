package client

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
)

// GithubClientContext provides a client and background context to use the Github API
func GithubClientContext(username string, password string) (*github.Client, context.Context) {
	if username == "" || password == "" {
		client, ctx := createEmptyBackGroundContext()
		return client, ctx
	}

	// Create transport for authentication
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	ctx := context.Background()
	user, resp, err := client.Users.Get(ctx, "")

	// If authentication fails, use an empty background context
	if err != nil || user == nil || resp.StatusCode == 401 {
		client, ctx := createEmptyBackGroundContext()
		return client, ctx
	}
	return client, ctx
}

func createEmptyBackGroundContext() (*github.Client, context.Context) {
	client := github.NewClient(nil)
	ctx := context.Background()
	return client, ctx
}
