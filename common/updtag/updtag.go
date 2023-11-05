package updtag

import (
	"context"

	"github.com/google/go-github/github"
)

func FetchTag() (string, error) {
	ctx := context.Background()

	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, "Florennum", "rudis")

	if err != nil {
		return "", err
	}

	tagName := *release.TagName
	return tagName, nil
}
