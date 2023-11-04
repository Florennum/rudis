package fvsetge

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Florennum/rudis/common/fetchtag"
)

func Fset() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(homeDir, ".var", "app", "io.github.vinegarhq.Vinegar", "config", "vinegar", "config.toml")

	var config map[string]interface{}
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatalf("Error decoding the config file: %v", err)
	}

	tag, err := fetchtag.FetchTag()
	if err != nil {
		log.Fatalf("Error fetching tag: %v", err)
	}

	extractedDir := filepath.Join(homeDir, ".local", "share", "rudis", "winege-ext")
	newValue := filepath.Join(extractedDir, fmt.Sprintf("lutris-%s-x86_64", tag))
	config["wineroot"] = newValue

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatalf("Error creating the config file: %v", err)
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(config); err != nil {
		log.Fatalf("Error encoding the config: %v", err)
	}
}
