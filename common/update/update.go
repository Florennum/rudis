package update

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Florennum/rudis/common/updtag"
)

func UpdateRudis() {
	latestTag, err := updtag.FetchTag()
	if err != nil {
		fmt.Println("Failed to fetch the latest 'rudis' tag:", err)
		return
	}

	currentTag := "v1.0.0-alpha"

	if latestTag != currentTag {
		fmt.Printf("Updating 'rudis' to version %s...\n", latestTag)

		projectPath := "./"
		err := os.Chdir(projectPath)
		if err != nil {
			fmt.Println("Failed to change working directory:", err)
			return
		}

		cmd := exec.Command("git", "pull")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Failed to run 'git pull':", err)
			return
		}
	} else {
		fmt.Println("No updates available at the moment.")
	}
}
