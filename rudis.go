package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dependcheck()
	vinegar := flag.Bool("vinegar", false, "vinegar")
	fvinegar := flag.Bool("f-vinegar", false, "flatpak vinegar")
	grapejuice := flag.Bool("grapejuice", false, "grapejuice")
	fgrapejuice := flag.Bool("f-grapejuice", false, "flatpak grapejuice")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: rudis [flag] [command]")
		return
	}

	command := flag.Arg(0)

	switch command {
	case "install-ge":
		installWINEGE(*vinegar, *fvinegar, *grapejuice, *fgrapejuice)

	case "update":
		update()
	case "patch-wayland":
		patchwayland(*vinegar, *fvinegar, *grapejuice, *fgrapejuice)
	case "help":
		Help()
	default:
		fmt.Println("Unknown command, use ./rudis help for help")
	}
}

func installWINEGE(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("Installing GE...")
	mkdir()
	setupwinege()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		showFailureNotification("Error installing GE: " + err.Error())
		return
	}

	if vinegar {
		configFile = filepath.Join(homeDir, ".config", "vinegar", "config.toml")
		vsetge()
	} else if fvinegar {
		patchflatpak()
		configFile = filepath.Join(homeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "vinegar", "config.toml")
		vsetge()
	} else if grapejuice {
		gjconfigFile = filepath.Join(homeDir, ".config", "brinkervii", "grapejuice", "user_settings.json")
		gjsetge()
	} else if fgrapejuice {
		patchflatpak()
		gjconfigFile = filepath.Join(homeDir, ".var", "app", "net.brinkervii.grapejuice", "config", "brinkervii", "grapejuice", "user_settings.json")
		gjsetge()
	} else {
		showFailureNotification("No installation option selected.")
	}
}

func patchwayland(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("ah boy here we go again")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		showFailureNotification("Error patching Wayland: " + err.Error())
		return
	}

	if vinegar {
		configFile = filepath.Join(homeDir, ".config", "vinegar", "config.toml")
		vpatchwayland()
		showSuccessNotification("Wayland patching for Vinegar successful!")
	} else if fvinegar {
		configFile = filepath.Join(homeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "vinegar", "config.toml")
		vpatchwayland()
		showSuccessNotification("Wayland patching for Flatpak Vinegar successful!")
	} else if grapejuice {
		showFailureNotification("Wayland patching for Grapejuice is currently not supported!")
	} else if fgrapejuice {
		showFailureNotification("Wayland patching for Flatpak Grapejuice is currently not supported!")
	} else {
		showFailureNotification("No patching option selected.")
	}
}
