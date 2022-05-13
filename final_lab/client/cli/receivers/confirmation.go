package receivers

import (
	"bufio"
	"drum-hero/cli/parsers"
	"drum-hero/models/commands"
	"fmt"
	"strings"
)

// Prompt for/receive confirmation input,
// parse it, and return parsed command to handler
func ReceiveConfirmation(r *bufio.Reader) commands.Confirmation {
	for {
		fmt.Print("(Y/N): ")

		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Unexpected EOF!")
			continue
		}

		input := strings.ToLower(strings.TrimSpace(raw))
		if len(input) < 1 {
			fmt.Println("No confirmation specified!")
			continue
		}

		return parsers.ParseConfirmation(input)
	}
}
