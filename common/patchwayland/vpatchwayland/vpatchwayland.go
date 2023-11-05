package vpatchwayland

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func UpdateLauncher() {
	var height, width, fps, unFocused int

	fmt.Println("To patch the Wayland cursor locking, we use gamescope here, so make sure it's actually installed.")
	fmt.Println("This will currently overwrite your launcher.")
	fmt.Println("Unfocused FPS caps your FPS while the game is unfocused.")

	fmt.Print("Type your screen height: ")
	_, err := fmt.Scan(&height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Type your screen width: ")
	_, err = fmt.Scan(&width)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Type what you want to limit your FPS to (put a big number for unlimited): ")
	_, err = fmt.Scan(&fps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Type what you want your unfocused FPS to be set to: ")
	_, err = fmt.Scan(&unFocused)
	if err != nil {
		log.Fatal(err)
	}

	// Construct the Gamescope launcher command
	gamescopeCommand := fmt.Sprintf("gamescope -h %d -w %d -W %d -H %d -r %d -o %d -f --force-grab-cursor --", height, width, height, width, fps, unFocused)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configFile := filepath.Join(homeDir, ".config", "vinegar", "config.toml")

	var config map[string]interface{}
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	// Ensure [player] section exists in the config
	playerSection, playerSectionExists := config["player"].(map[string]interface{})
	if !playerSectionExists {
		playerSection = make(map[string]interface{})
		config["player"] = playerSection
	}

	playerSection["launcher"] = gamescopeCommand

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(config); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully updated launcher!")
}
