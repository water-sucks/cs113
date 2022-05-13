package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Enum representation for possible speeds
type Difficulty int

const (
	Easy Difficulty = iota
	Normal
	Hard
	InvalidDifficulty
)

func (d Difficulty) ToString() string {
	return []string{"Easy", "Normal", "Hard"}[d]
}

type Settings struct {
	Difficulty Difficulty `json:"difficulty"`
	Measures   int        `json:"measures"`
}

var (
	configDirectory = os.ExpandEnv("$HOME/.config/drum-hero")
	configLocation  = configDirectory + "/settings.json"
	defaultSettings = Settings{
		Difficulty: Normal,
		Measures:   4,
	}
)

// Read settings in from JSON configuration file
// located at $HOME/.config/drum-hero/settings.json.
// This will only work on a Linux system or a system
// compliant with XDG specifications, so it will not
// work on Windows. This returns defaults if an error
// is encountered.
func ReadSettings() (Settings, error) {
	file, err := os.Open(configLocation)
	if err != nil {
		fmt.Print("Unable to open settings file, proceeding with defaults\n\n")
		return defaultSettings, err
	}
	defer func() { _ = file.Close() }()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print("Could not read configuration file data, overwriting with defaults\n\n")
		if err != nil {
			return Settings{}, err
		}
		return defaultSettings, err
	}

	var s Settings

	err = json.Unmarshal(raw, &s)
	if err != nil {
		fmt.Print("Could not parse configuration file data, overwriting with defaults\n\n")
		return defaultSettings, err
	}

	if err := ValidateSettings(s); err != nil {
		fmt.Printf("Error: %s, overwriting with defaults\n\n", err.Error())
		return defaultSettings, err
	}

	return s, nil
}

// Write specified settings to configuration file.
func PersistSettings(s Settings) error {
	contents, _ := json.MarshalIndent(s, "", "  ")

	if _, err := os.Stat(configDirectory); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(configDirectory, 0755)
		if err != nil {
			fmt.Println("Unable to create config directory: ", err.Error())
			return err
		}
	}
	err := ioutil.WriteFile(configLocation, contents, 0644)
	if err != nil {
		fmt.Println("Unable to write settings to configuration file: ", err.Error())
		return err
	}

	return nil
}

// Validate given settings and make sure they make sense.
func ValidateSettings(s Settings) error {
	if s.Difficulty < 0 || s.Difficulty > 2 {
		return errors.New("invalid difficulty level, must be in range 0-2")
	}

	if s.Measures < 1 {
		return errors.New("measures must be greater than zero")
	}

	return nil
}
