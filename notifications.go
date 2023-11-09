package main

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func showSuccessNotification(message string) {
	err := beeep.Notify("Installation Notification", message, "path/to/success/icon")
	if err != nil {
		fmt.Println("Error displaying success notification:", err)
	}
}

func showFailureNotification(message string) {
	err := beeep.Alert("Installation Error", message, "path/to/failure/icon")
	if err != nil {
		fmt.Println("Error displaying failure notification:", err)
	}
}

func showDependencyFailureNotification(message string) {
	err := beeep.Alert("Dependency Error", message, "path/to/dependency/icon")
	if err != nil {
		fmt.Println("Error displaying dependency failure notification:", err)
	}
}
