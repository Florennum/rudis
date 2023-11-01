package main

import (
	"fmt"
	"os"

	"github.com/Florennum/rudis/common/downloadge"
	"github.com/Florennum/rudis/common/mkdir"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myprogram [command]")
		return
	}

	command := os.Args[1]

	switch command {
	case "install-ge":
		fmt.Println("Installing GE...")
		mkdir.Mkdir()
		downloadge.Downloadge()

	case "update":
		fmt.Println("Updating something...")
		// Call a function to perform an update.

	default:
		fmt.Println("Unknown command:", command)
	}
}
