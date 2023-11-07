package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
)

func updtag() (string, error) {
	ctx := context.Background()

	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, "Florennum", "rudis")

	if err != nil {
		return "", err
	}

	tagName := *release.TagName
	return tagName, nil
}

func update() error {
	latestTag, err := updtag()
	if err != nil {
		return fmt.Errorf("failed to fetch the latest 'rudis' tag: %v", err)
	}

	currentTag := "v2.0.0-alpha"

	if latestTag != currentTag {
		fmt.Printf("Updating rudis to version %s...\n", latestTag)

		projectPath := "./"
		err = os.Chdir(projectPath)
		if err != nil {
			return fmt.Errorf("failed to change working directory: %v", err)
		}

		cmd := exec.Command("git", "pull")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run 'git pull': %v", err)
		}
	} else {
		fmt.Println("No updates available at the moment.")
	}

	return nil
}
