package vsetge

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Florennum/rudis/common/fetchtag"
	tomlv2 "github.com/pelletier/go-toml/v2"
)

func ensureWineRootKey(config map[string]interface{}) {
	// Check if the 'wineroot' key exists in the config
	_, ok := config["wineroot"]
	if !ok {
		// 'wineroot' key is missing, create it
		config["wineroot"] = ""
	}
}

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
	// Define the path to the TOML configuration file
	configFile := filepath.Join(currentUser.HomeDir, ".config", "vinegar", "config.toml")

	// Read the contents of the TOML configuration file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]interface{}
	ensureWineRootKey(config)

	// Unmarshal the TOML content into a Go data structure
	err = tomlv2.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Update the 'wineroot' value in the TOML data structure
	extractedDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
	newValue := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))
	config["wineroot"] = newValue

	// Marshal the updated data back to TOML format
	updatedTOML, err := tomlv2.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated configuration back to the file
	err = os.WriteFile(configFile, updatedTOML, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated 'wineroot' in the configuration file:", configFile)
}

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

	// Define the path to the TOML configuration file
	configFile := filepath.Join(currentUser.HomeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "config.toml")

	// Read the contents of the TOML configuration file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]interface{}

	// Unmarshal the TOML content into a Go data structure
	err = tomlv2.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}
	ensureWineRootKey(config)
	// Update the 'wineroot' value in the TOML data structure
	extractedDir := filepath.Join(currentUser.HomeDir, ".local", "share", "rudis", "winege-ext")
	newValue := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))
	config["wineroot"] = newValue

	// Marshal the updated data back to TOML format
	updatedTOML, err := tomlv2.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated configuration back to the file
	err = os.WriteFile(configFile, updatedTOML, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated 'wineroot' in the configuration file:", configFile)
}
