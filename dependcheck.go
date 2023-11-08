package main

import (
	"fmt"
	"os"
	"os/exec"
)

func checkDependency(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func dependcheck() {
	dependencies := []string{"axel", "pv", "tar"}
	missingCount := 0

	for _, dep := range dependencies {
		if !checkDependency(dep) {
			fmt.Printf("%s is not installed!\n", dep)
			missingCount++
		}
	}

	if missingCount == 1 {
		fmt.Printf("%d dependency is missing\n", missingCount)
		os.Exit(1)
	} else if missingCount > 1 {
		fmt.Printf("%d dependencies are missing\n", missingCount)
		os.Exit(1)
	}
}
