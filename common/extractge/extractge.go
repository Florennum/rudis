package extractge

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
)

func ExtractGE() {
	tag, err := fetchtag.FetchTag()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Define the target directory for extraction
	targetDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-arch")
	archivePath := filepath.Join(targetDir, fmt.Sprintf("wine-lutris-%s-x86_64.tar.xz", tag))

	// Define the output directory
	outputDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")

	// Check if the output directory exists
	_, outputDirErr := os.Stat(outputDir)
	if outputDirErr == nil {
		fmt.Println("Archive is already extracted in:", outputDir)
		return
	}

	// Create the output directory
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("pv %s | tar -xJf - -C %s", archivePath, outputDir))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error extracting archive:", err)
		return
	}

	fmt.Println("Archive extracted successfully to:", outputDir)
}
