package handlers

import (
	"bufio"
	receivers2 "drum-hero/cli/receivers"
	"drum-hero/models"
	"drum-hero/models/commands"
	"fmt"
)

func SettingsHandler(r *bufio.Reader, s *models.Settings) {
	temporarySettings := *s
	saved := true

	ViewSettingsHandler(s, true, true)
	fmt.Println()
	SettingsHelp()

main:
	for {
		command := receivers2.ReceiveSettingsInput(r)

		switch command {
		case commands.ChangeSetting:
			ChangeSettingsHandler(r, &temporarySettings, &saved)
		case commands.ViewSettings:
			ViewSettingsHandler(&temporarySettings, saved, true)
		case commands.SaveSettings:
			err := models.ValidateSettings(temporarySettings)
			if err != nil {
				fmt.Println("Unable to validate settings, exiting settings menu")
				break main
			}

			*s = temporarySettings
			saved = true
			err = models.PersistSettings(*s)
			if err != nil {
				fmt.Println("Unable to save configuration file")
				break main
			}
			fmt.Println("Saved settings")
			break main
		case commands.SettingsHelp:
			SettingsHelp()
		case commands.QuitSettings:
			var confirmation commands.Confirmation
			if temporarySettings != *s {
			confirmation:
				for {
					fmt.Print("There are unsaved settings! Are you sure? ")

					confirmation = receivers2.ReceiveConfirmation(r)

					if confirmation != commands.Yes {
						if confirmation != commands.No {
							fmt.Println("Unable to receive confirmation, retrying")
							continue
						}
						break confirmation
					}
					fmt.Println("Exiting without saving settings")
					break main
				}
			}
			if confirmation == commands.No {
				continue main
			}
			fmt.Println("Exiting to main menu")
			break main
		}
	}
}

func ChangeSettingsHandler(r *bufio.Reader, s *models.Settings, saved *bool) {
main:
	for {
		command := receivers2.ReceiveChangeSettingsInput(r)

		switch command {
		case commands.ChangeDifficulty:
			if changed := ChangeDifficultyHandler(r, s); changed {
				fmt.Printf("Changed difficulty to %s\n", s.Difficulty.ToString())
			}
			*saved = false
		case commands.ChangeMeasures:
			if changed := ChangeMeasuresHandler(r, s); changed {
				fmt.Printf("Changed difficulty to %d\n", s.Measures)
			}
			*saved = false
		case commands.QuitChangeSettings:
			fmt.Println("Going back to settings menu")
			break main
		}
	}
}

func ChangeDifficultyHandler(r *bufio.Reader, s *models.Settings) bool {
	for {
		command := receivers2.ReceiveChangeDifficultyInput(r)

		if command != models.InvalidDifficulty {
			s.Difficulty = command
			return true
		}

		return false
	}
}

func ChangeMeasuresHandler(r *bufio.Reader, s *models.Settings) bool {
	for {
		measures := receivers2.ReceiveChangeMeasuresInput(r)
		s.Measures = measures
		return true
	}
}

func ViewSettingsHandler(s *models.Settings, saved bool, displaySavedStatus bool) {
	fmt.Println("Here are your current settings:")
	fmt.Printf("Difficulty: %s\n", s.Difficulty.ToString())
	fmt.Printf("Measures: %d\n", s.Measures)
	if displaySavedStatus {
		if !saved {
			fmt.Println("These settings have not been saved yet.")
		}
	}
}
