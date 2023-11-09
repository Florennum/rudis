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
		fmt.Println("Error getting home directory:", err)
		showFailureNotification("Error getting home directory: " + err.Error())
		return
	}

	tag, err := FetchTag()
	if err != nil {
		showFailureNotification("Error fetching tag: " + err.Error())
		return
	}

	extractedDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-ext")
	dirpath := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))

	cmd := exec.Command("flatpak", "override", "--user", "--filesystem="+dirpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error patching flatpak:", err)
		showFailureNotification("Error patching flatpak: " + err.Error())
		return
	}

	fmt.Println("Successfully patched flatpak to be able to use:", dirpath)
}
