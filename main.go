package main

import (
	"flag"
	"fmt"

	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/extractge"
	"github.com/Florennum/rudis/common/mkdir"
	"github.com/Florennum/rudis/common/setge"
)

func main() {
	vinegar := flag.Bool("vinegar", false, "vinegar")
	// fvinegar := flag.Bool("f-vinegar", false, "flatpak vinegar")
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
			setge.Vsetge()
		}

	case "update":
		fmt.Println("Updating something...")

	default:
		fmt.Println("Unknown command:", command)
	}
}
