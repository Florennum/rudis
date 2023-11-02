package gjsetge

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
)

func Setge() {
	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	tag, err := fetchtag.FetchTag()
	if err != nil {
		return
	}

	// Define the path to the JSON configuration file for Grapejuice
	configFile := filepath.Join(currentUser.HomeDir, ".config", "brinkervii", "grapejuice", "user_settings.json")

	// Read the contents of the JSON configuration file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]string

	// Unmarshal the JSON content into a Go data structure
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Update "wine_home" and "default_wine_home" values
	extractedDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
	newValue := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))
	config["wine_home"] = newValue
	config["default_wine_home"] = newValue

	// Marshal the updated data back to JSON format
	updatedJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated configuration back to the file
	err = os.WriteFile(configFile, updatedJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated 'wine_home' and 'default_wine_home' in the JSON configuration file:", configFile)
}

// /.var/app/net.brinkervii.grapejuice/config/brinkervii/grapejuice/user_settings.json

func Fsetge() {
	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	tag, err := fetchtag.FetchTag()
	if err != nil {
		return
	}

	// what in the actual fuck is this brinker
	configFile := filepath.Join(currentUser.HomeDir, ".var", "app", "net.brinkervii.grapejuice", "config", "brinkervii", "grapejuice", "user_settings.json")

	// Read the contents of the JSON configuration file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]string

	// Unmarshal the JSON content into a Go data structure
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Update "wine_home" and "default_wine_home" values
	extractedDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
	newValue := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))
	config["wine_home"] = newValue
	config["default_wine_home"] = newValue

	// Marshal the updated data back to JSON format
	updatedJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated configuration back to the file
	err = os.WriteFile(configFile, updatedJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated 'wine_home' and 'default_wine_home' in the JSON configuration file:", configFile)
}
