package main

import (
	"fmt"
	"os"

	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/extractge"
	"github.com/Florennum/rudis/common/mkdir"
	"github.com/Florennum/rudis/common/setge"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rudis [command]")
		return
	}

	command := os.Args[1]

	switch command {
	case "install-ge":
		fmt.Println("Installing GE...")
		mkdir.Mkdir()
		downloadge.Downloadge()
		extractge.ExtractGE()
		setge.Vsetge()

	case "update":
		fmt.Println("Updating something...")
		// Call a function to perform an update.

	default:
		fmt.Println("Unknown command:", command)
	}
}
