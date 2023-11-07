package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func patchflatpak() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tag, err := FetchTag()
	if err != nil {
		return
	}

	extractedDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-ext")
	dirpath := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))

	cmd := exec.Command("flatpak", "override", "--user", "--filesystem="+dirpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("There was an error while trying to patch flatpak:", err)
		return
	}

	fmt.Println("Successfully patched flatpak to be able to use:", dirpath)
}
