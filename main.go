package main

import (
	"flag"
	"fmt"

	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/extractge"
	"github.com/Florennum/rudis/common/mkdir"
	"github.com/Florennum/rudis/common/patchflatpak"
	"github.com/Florennum/rudis/common/setge/gjsetge"
	"github.com/Florennum/rudis/common/setge/vsetge"
)

func main() {
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
		fmt.Println("Installing GE...")
		mkdir.Mkdir()
		downloadge.Downloadge()
		extractge.ExtractGE()

		if *vinegar {
			vsetge.Setge()
		}
		if *fvinegar {
			patchflatpak.Patch()
			vsetge.Fsetge()
		}
		if *grapejuice {
			gjsetge.Setge()
		}
		if *fgrapejuice {
			patchflatpak.Patch()
			gjsetge.Fsetge()
		}

	case "update":
		fmt.Println("Updating something...")

	default:
		fmt.Println("Unknown command:", command)
	}
}
