package patchflatpak

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
)

func Patch() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	tag, err := fetchtag.FetchTag()
	if err != nil {
		return
	}

	extractedDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
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
