package parsers

import (
	"drum-hero/models/commands"
)

// Command parser for main shell
func ParseMainCommand(input string) commands.MainCommand {
	switch input {
	case "start":
		return commands.Start
	case "settings":
		return commands.Settings
	case "instructions":
		return commands.Instructions
	case "help":
		return commands.Help
	case "quit":
		return commands.Quit
	}

	return commands.InvalidMainCommand
}
