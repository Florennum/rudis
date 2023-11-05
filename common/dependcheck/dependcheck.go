package dependcheck

import (
	"fmt"
	"os/exec"
)

func checkDependency(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func Check() {
	dependencies := []string{"axel", "pv", "tar"}
	missingCount := 0

	for _, dep := range dependencies {
		if !checkDependency(dep) {
			fmt.Printf("%s is not installed\n", dep)
			missingCount++
		}
	}

	if missingCount > 0 {
		fmt.Printf("%d dependencies are missing\n", missingCount)
	} else {
		fmt.Println("All dependencies are installed")
	}
}
