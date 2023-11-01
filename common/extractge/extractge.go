package extractge

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
)

func ExtractGEArchive() error {
	tag, err := fetchtag.FetchTag()
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	// Define the target directory for extraction
	targetDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
	archivePath := filepath.Join(targetDir, fmt.Sprintf("wine-lutris-%s-x86_64.tar.xz", tag))

	// Open the ZIP archive
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	// Extract files from the ZIP archive
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the output file
		outputFile, err := os.Create(filepath.Join(targetDir, f.Name))
		if err != nil {
			return err
		}
		defer outputFile.Close()

		// Copy the contents to the output file
		_, err = io.Copy(outputFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
