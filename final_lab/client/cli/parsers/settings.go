package parsers

import (
	"drum-hero/models"
	"drum-hero/models/commands"
	"strconv"
)

// Main settings shell parser
func ParseSettingsCommand(input string) commands.SettingsCommand {
	switch input {
	case "change":
		return commands.ChangeSetting
	case "view":
		return commands.ViewSettings
	case "save":
		return commands.SaveSettings
	case "help":
		return commands.SettingsHelp
	case "quit":
		return commands.QuitSettings
	}

	return commands.InvalidSettingsCommand
}

// Convert number to numbered constant and reference
// constant based on it (this is based on the `iota`
// construct in Go, which are typed integer constants)
func ParseChangeSettingCommand(input string) commands.ChangeSettingCommand {
	number, err := strconv.Atoi(input)
	if err != nil {
		return commands.InvalidChangeSettingCommand
	}
	number -= 1

	if number < 0 && number > 3 {
		return commands.InvalidChangeSettingCommand
	}

	return commands.ChangeSettingCommand(number)
}

// Get difficulty level from user
func ParseDifficulty(input string) models.Difficulty {
	switch input {
	case "easy":
		return models.Easy
	case "normal":
		return models.Normal
	case "hard":
		return models.Hard
	}

	return models.InvalidDifficulty
}

// Get number of measures ;-1 represents invalid input
func ParseMeasures(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil || number < 1 {
		return -1
	}

	return number
}
