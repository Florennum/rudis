package main

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: rudis --flag argument")
	fmt.Println("Available flags: --vinegar --f-vinegar --grapejuice --f-grapejuice")
	fmt.Println("Available arguments: install-ge patch-wayland")
	fmt.Println("Example: ./rudis --vinegar install-ge")
	fmt.Println("Wayland patcher for grapejuice is currently unsupported!")
}
