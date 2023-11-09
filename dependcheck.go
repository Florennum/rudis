// dependcheck.go
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
	missingDependencies := ""

	for _, dep := range dependencies {
		if !checkDependency(dep) {
			fmt.Printf("%s is not installed!\n", dep)
			missingDependencies += dep + "\n"
			missingCount++
		}
	}

	if missingCount == 1 {
		showDependencyFailureNotification(fmt.Sprintf("%d dependency is missing:\n%s", missingCount, missingDependencies))
		os.Exit(1)
	} else if missingCount > 1 {
		showDependencyFailureNotification(fmt.Sprintf("%d dependencies are missing:\n%s", missingCount, missingDependencies))
		os.Exit(1)
	}
}
