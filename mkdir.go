package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func mkdir() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		showFailureNotification("Error getting current user: " + err.Error())
		return
	}

	baseDir := filepath.Join(currentUser.HomeDir, ".local", "share")
	subDirs := []string{"rudis", "rudis/winege-arch", "rudis/winege-ext"}

	for _, subDir := range subDirs {
		dirPath := filepath.Join(baseDir, subDir)

		_, err := os.Stat(dirPath)
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				fmt.Println("Error creating directory:", err)
				showFailureNotification("Error creating directory: " + err.Error())
				return
			}
		} else if err != nil {
			fmt.Println("Error checking directory:", err)
			showFailureNotification("Error checking directory: " + err.Error())
			return
		}
	}
}
