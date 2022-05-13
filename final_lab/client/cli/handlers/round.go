package handlers

import (
	"bufio"
	"drum-hero/cli/receivers"
	"drum-hero/models"
	"drum-hero/models/commands"
	"fmt"
)

func RoundHandler(r *bufio.Reader, s models.Settings) {
	var round models.Round

	// Create patterns and shows them until
	// user selects one they like
	for {
		round = models.CreateRound(s)

		round.Pattern.PrettyPrint()

		fmt.Println("Use this pattern? ")
		if confirmation := receivers.ReceiveConfirmation(r); confirmation == commands.Yes {
			break
		}
	}

	StartInstructions(s.Difficulty)

	// Confirm that user wants to start
	fmt.Print("Are you ready? ")
	if confirmation := receivers.ReceiveConfirmation(r); confirmation != commands.Yes {
		fmt.Println("No confirmation given, exiting back to main menu")
		return
	} else {
		fmt.Print("Hit the pad whenever you're ready to start the round.\n\n")
	}

	// Receive parsed hits in form of Pattern
	pattern, err := receivers.ReceiveHits(s)
	if err != nil {
		fmt.Println("There was an error: ", err.Error())
		return
	}

	// Show what was played and the original pattern as well
	// so that user knows what was scanner
	fmt.Println("Original pattern: ")
	round.Pattern.PrettyPrint()

	fmt.Println("What you played:")
	pattern.PrettyPrint()

	// Evaluate truth expression to check if user has
	// played the pattern correctly, and show appropriate
	// message for if the user has won or lost
	if round.Evaluator(pattern) {
		fmt.Println("You won! :)")
	} else {
		fmt.Println("You lost :( better luck next time!")
	}
}
