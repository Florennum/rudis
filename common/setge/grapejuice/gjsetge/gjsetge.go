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

func Set() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	tag, err := fetchtag.FetchTag()
	if err != nil {
		log.Fatal(err)
	}

	// Replace %s with the tag in the paths
	defaultWineHome := fmt.Sprintf("%s/.local/share/rudis/winege-ext/lutris-%s-x86_64/", homeDir, tag)
	wineHome := fmt.Sprintf("%s/.local/share/rudis/winege-ext/lutris-%s-x86_64/", homeDir, tag)

	configFile := filepath.Join(currentUser.HomeDir, ".config", "brinkervii", "grapejuice", "user_settings.json")

	// Read the JSON file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a struct to represent the specific fields to update
	type UpdateData struct {
		DefaultWineHome string `json:"default_wine_home"`
		WineHome        string `json:"wine_home"`
	}

	updateData := UpdateData{
		DefaultWineHome: defaultWineHome,
		WineHome:        wineHome,
	}

	// Unmarshal the JSON content into a Go data structure
	var config map[string]interface{}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Update the specific field(s) in the JSON data structure
	config["default_wine_home"] = updateData.DefaultWineHome
	config["wine_home"] = updateData.WineHome

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

	fmt.Printf("Updated 'default_wine_home' and 'wine_home' in the JSON configuration file: %s\n", configFile)
}
