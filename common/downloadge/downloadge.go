package downloadge

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
)

func Downloadge() error {
	tag, err := fetchtag.FetchTag()
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	targetDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-arch")

	// Specify the output file path
	outputFile := filepath.Join(targetDir, fmt.Sprintf("wine-lutris-%s-x86_64.tar.xz", tag))

	// Check if the file already exists
	if _, fileErr := os.Stat(outputFile); fileErr == nil {
		fmt.Printf("Archive already downloaded: %s\n", outputFile)
	}

	// Build the download URL
	downloadURL := fmt.Sprintf("https://github.com/GloriousEggroll/wine-ge-custom/releases/download/%s/wine-lutris-%s-x86_64.tar.xz", tag, tag)

	cmd := exec.Command("axel", "-o", outputFile, downloadURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}
