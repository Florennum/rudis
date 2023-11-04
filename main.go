package main

import (
	"flag"
	"fmt"

	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/extractge"
	"github.com/Florennum/rudis/common/mkdir"
	"github.com/Florennum/rudis/common/patchflatpak"
	"github.com/Florennum/rudis/common/setge/vsetge"
)

func main() {
	// Define command-line flags
	vinegar := flag.Bool("vinegar", false, "vinegar")
	fvinegar := flag.Bool("f-vinegar", false, "flatpak vinegar")
	grapejuice := flag.Bool("grapejuice", false, "grapejuice")
	fgrapejuice := flag.Bool("f-grapejuice", false, "flatpak grapejuice")

	// Parse command-line arguments
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
		updateSomething()

	default:
		fmt.Println("Unknown command:", command)
	}
}

func installGE(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("Installing GE...")

	// Create a directory
	mkdir.Mkdir()

	// Download files
	downloadge.Downloadge()

	// Extract files from an archive
	extractge.ExtractGE()

	if vinegar {
		vsetge.Set()
	}
	if fvinegar {
		patchflatpak.Patch()
		vsetge.Fset()
	}
	if grapejuice {
		// setge.GJsetge()
	}
	if fgrapejuice {
		patchflatpak.Patch()
		// setge.FGJsetge()
	}
}

func updateSomething() {
	fmt.Println("Updating something...")
}
