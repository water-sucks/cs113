package receivers

import (
	"bufio"
	"drum-hero/cli/parsers"
	"drum-hero/models/commands"
	"fmt"
	"strings"
)

// Prompt for/receive main command input,
// parse it, and return command back to handler
func ReceiveMainInput(r *bufio.Reader) commands.MainCommand {
	for {
		fmt.Print(">>> ")

		rawInput, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("\nUnexpected EOF")
			continue
		}

		input := strings.ToLower(strings.TrimSpace(rawInput))

		if parsed := parsers.ParseMainCommand(input); parsed == commands.InvalidMainCommand {
			fmt.Println(`Invalid command! Type "help" for a list of commands.`)
		} else {
			return parsed
		}
	}
}
