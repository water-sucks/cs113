package receivers

import (
	"bufio"
	"drum-hero/cli/parsers"
	"drum-hero/models"
	"drum-hero/models/commands"
	"fmt"
	"strings"
)

// Prompt for/receive main command input,
// parse it, and return command back to
// main settings handler to decide which
// sub-handler to run
func ReceiveSettingsInput(r *bufio.Reader) commands.SettingsCommand {
	for {
		fmt.Print("(settings): ")

		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Unexpected EOF!")
			continue
		}

		input := strings.TrimSpace(raw)
		if len(input) < 1 {
			fmt.Println("No command specified")
			continue
		}

		if parsed := parsers.ParseSettingsCommand(input); parsed == commands.InvalidSettingsCommand {
			fmt.Println(`Invalid command! Type "help" for a list of commands.`)
		} else {
			return parsed
		}
	}
}

// Prompt for/receive which setting to change,
// parse using `iota` conversion, and return
// command to execute based on that
func ReceiveChangeSettingsInput(r *bufio.Reader) commands.ChangeSettingCommand {
	fmt.Print(`Change:
  1) Difficulty
  2) Measures
  3) Return to settings menu
Choose option (1-3): `)
	for {
		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Unexpected EOF!")
			continue
		}

		input := strings.TrimSpace(raw)
		if len(input) < 1 {
			fmt.Print("Enter 1-2 to change a setting, or 3 to return back.")
			continue
		}

		if parsed := parsers.ParseChangeSettingCommand(input); parsed == commands.InvalidChangeSettingCommand {
			fmt.Println("Enter 1-2 to change a setting, or 3 to return back.")
			continue
		} else {
			return parsed
		}
	}
}

// Prompt for/receive a valid Difficulty input
// and return back to handler to use
func ReceiveChangeDifficultyInput(r *bufio.Reader) models.Difficulty {
	for {
		fmt.Print("Choose difficulty (easy, normal, or hard): ")

		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Unexpected EOF!")
			continue
		}

		input := strings.ToLower(strings.TrimSpace(raw))
		if len(input) < 1 {
			fmt.Println("Difficulty must be easy, normal, or hard.")
			continue
		}

		if parsed := parsers.ParseDifficulty(input); parsed == models.InvalidDifficulty {
			fmt.Println("Difficulty must be easy, normal or hard.")
		} else {
			return parsed
		}
	}
}

// Prompt for/receive a valid number of measures
// to play and return back to handler to use
func ReceiveChangeMeasuresInput(r *bufio.Reader) int {
	for {
		fmt.Print("Choose number of measures to play (at least 1): ")

		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Unexpected EOF!")
			continue
		}

		input := strings.TrimSpace(raw)
		if len(input) < 1 {
			fmt.Println("No number of measures specified")
			continue
		}

		if parsed := parsers.ParseMeasures(input); parsed < 0 {
			fmt.Println("Measures must be a number greater than zero.")
		} else {
			return parsed
		}
	}
}
