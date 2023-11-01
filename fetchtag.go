package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

func FetchTag() {
	ctx := context.Background()

	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, "GloriousEggroll", "wine-ge-custom")

	if err != nil {
		fmt.Printf("There was a problem fetching the latest release: %v\n", err)
		os.Exit(1)
	}

	tagName := *release.TagName
	fmt.Printf("%s", tagName)
}
