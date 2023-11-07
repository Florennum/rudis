package main

import (
	"context"

	"github.com/google/go-github/github"
)

func FetchTag() (string, error) {
	ctx := context.Background()

	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, "GloriousEggroll", "wine-ge-custom")

	if err != nil {
		return "", err
	}

	tagName := *release.TagName
	return tagName, nil
}
