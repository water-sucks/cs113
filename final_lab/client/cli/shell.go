package cli

import (
	"bufio"
	"drum-hero/cli/handlers"
	"drum-hero/cli/receivers"
	"drum-hero/models"
	"drum-hero/models/commands"
	"fmt"
	"os"
)

func Shell() error {
	// Single reader from stdin for all text input
	// to reduce latency
	reader := bufio.NewReader(os.Stdin)

	// Read settings from configuration file
	// or create if nonexistent
	settings, err := models.ReadSettings()
	if err != nil {
		_ = models.PersistSettings(settings)
	}

	handlers.Instructions()
	fmt.Println()
	handlers.MainHelp()
main:
	for {
		command := receivers.ReceiveMainInput(reader)

		switch command {
		case commands.Start:
			handlers.RoundHandler(reader, settings)
		case commands.Settings:
			handlers.SettingsHandler(reader, &settings)
		case commands.Instructions:
			handlers.Instructions()
		case commands.Help:
			handlers.MainHelp()
		case commands.Quit:
			break main
		}
	}
	fmt.Println("Thanks for playing Drum Hero! :)")
	return nil
}
