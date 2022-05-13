package parsers

import (
	"drum-hero/models/commands"
	"strings"
)

// Common confirmation dialog that takes Y/N and
// returns a type-safe representation based on it.
func ParseConfirmation(input string) commands.Confirmation {
	letter := strings.ToLower(input)[0]

	switch letter {
	case 'y':
		return commands.Yes
	case 'n':
		return commands.No
	}

	return commands.InvalidConfirmation
}
