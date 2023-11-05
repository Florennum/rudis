package main

import (
	"flag"
	"fmt"

	"github.com/Florennum/rudis/common/dependcheck"
	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/extractge"
	"github.com/Florennum/rudis/common/mkdir"
	"github.com/Florennum/rudis/common/patchflatpak"
	"github.com/Florennum/rudis/common/patchwayland/fvpatchwayland"
	"github.com/Florennum/rudis/common/patchwayland/vpatchwayland"
	"github.com/Florennum/rudis/common/setge/grapejuice/fgjsetge"
	"github.com/Florennum/rudis/common/setge/grapejuice/gjsetge"
	"github.com/Florennum/rudis/common/setge/vinegar/fvsetge"
	"github.com/Florennum/rudis/common/setge/vinegar/vsetge"
	"github.com/Florennum/rudis/common/update"
)

func main() {
	dependcheck.Check()
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
		update.UpdateRudis()
	case "patch-wayland":
		patchwayland(*vinegar, *fvinegar, *grapejuice, *fgrapejuice)
	default:
		fmt.Println("Unknown command:", command)
	}
}

func installGE(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("Installing GE...")
	mkdir.Mkdir()
	downloadge.Downloadge()
	extractge.ExtractGE()

	if vinegar {
		vsetge.Set()
	}
	if fvinegar {
		patchflatpak.Patch()
		fvsetge.Fset()
	}
	if grapejuice {
		gjsetge.Set()
	}
	if fgrapejuice {
		patchflatpak.Patch()
		fgjsetge.Set()
	}
}

func patchwayland(vinegar, fvinegar, grapejuice, fgrapejuice bool) {
	fmt.Println("ah boy here we go again")

	if vinegar {
		vpatchwayland.UpdateLauncher()
	}

	if fvinegar {
		fvpatchwayland.UpdateLauncher()
	}
}
