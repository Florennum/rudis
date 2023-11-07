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
		installGE(*vinegar, *fvinegar, *grapejuice, *fgrapejuice)

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

func installGE(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("Installing GE...")
	mkdir()
	setupge()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if vinegar {
		configFile = filepath.Join(homeDir, ".config", "vinegar", "config.toml")
		vsetge()
	}
	if fvinegar {
		patchflatpak()
		configFile = filepath.Join(homeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "vinegar", "config.toml")
		vsetge()
	}
	if grapejuice {
		gjconfigFile = filepath.Join(homeDir, ".config", "brinkervii", "grapejuice", "user_settings.json")
		gjsetge()
	}
	if fgrapejuice {
		patchflatpak()
		gjconfigFile = filepath.Join(homeDir, ".config", "brinkervii", "grapejuice", "user_settings.json")
	}
}

func patchwayland(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("ah boy here we go again")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if vinegar {
		configFile = filepath.Join(homeDir, ".config", "vinegar", "config.toml")
		vpatchwayland()
	}

	if fvinegar {
		configFile = filepath.Join(homeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "vinegar", "config.toml")
		vpatchwayland()
	}

	if grapejuice {
		fmt.Println("This action is curreuntly not supported!")
	}

	if fgrapejuice {
		fmt.Println("This action is curreuntly not supported!")
	}
}
