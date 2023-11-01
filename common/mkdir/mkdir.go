package mkdir

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func Mkdir() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
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
				return
			}
			fmt.Println("Directory created successfully:", dirPath)
		} else if err != nil {
			fmt.Println("Error checking directory:", err)
			return
		}
	}
}
