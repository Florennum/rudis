package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func setupwinege() error {
	tag, err := FetchTag()
	if err != nil {
		showFailureNotification("Error fetching tag: " + err.Error())
		fmt.Println("Error fetching tag:", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		showFailureNotification("Error getting home directory: " + err.Error())
		fmt.Println("Error getting home directory:", err)
		return err
	}

	downloadDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-arch")

	// Specify the output file path
	outputFile := filepath.Join(downloadDir, fmt.Sprintf("wine-lutris-%s-x86_64.tar.xz", tag))

	// Check if the file already exists
	if _, fileErr := os.Stat(outputFile); fileErr == nil {
		fmt.Println("Archive already downloaded:", outputFile)
		return nil
	}

	// Build the download URL
	downloadURL := fmt.Sprintf("https://github.com/GloriousEggroll/wine-ge-custom/releases/download/%s/wine-lutris-%s-x86_64.tar.xz", tag, tag)

	cmd := exec.Command("axel", "-o", outputFile, downloadURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		showFailureNotification("Error downloading archive: " + err.Error())
		fmt.Println("Error downloading archive:", err)
		return fmt.Errorf("error: %v", err)
	}

	outputDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-ext")
	extractedDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-ext")
	check := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))

	// Check if the output directory exists
	_, outputDirErr := os.Stat(check)
	if outputDirErr == nil {
		fmt.Println("Archive is already extracted in:", check)
		return nil
	}

	// Create the output directory
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		showFailureNotification("Error creating output directory: " + err.Error())
		fmt.Println("Error creating output directory:", err)
		return err
	}

	fmt.Println("Archive extracted successfully to:", outputDir)

	cmd = exec.Command("sh", "-c", fmt.Sprintf("pv %s | tar -xJf - -C %s", outputFile, extractedDir))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		showFailureNotification("Error extracting archive: " + err.Error())
		fmt.Println("Error extracting archive:", err)
		return err
	}

	return nil
}