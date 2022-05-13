package receivers

import (
	"bufio"
	"context"
	"drum-hero/cli/parsers"
	"drum-hero/models"
	"fmt"
	"runtime"
	"time"

	"github.com/tarm/serial"
)

// Receive sensor input from Arduino based on
// user-provided settings. This only works on
// Linux (and likely only on my laptop) because
// of the way used to reference the device name.
func ReceiveHits(settings models.Settings) (models.Pattern, error) {
	// Get handle on serial port
	s, err := serial.OpenPort(&serial.Config{
		Name: "/dev/ttyACM0",
		Baud: 9600,
	})
	if err != nil {
		return nil, err
	}
	defer s.Close() // Close even if interrupted

	scanner := bufio.NewScanner(s)

	bpm := settings.Difficulty.ToBPM()                // Speed to play at
	duration := bpm.ToDuration(settings.Measures * 4) // Time to scan input for
	beats := settings.Measures * 8                    // Create Pattern input with size `beats`

	// Block thread until user hits pad to start.
	confirmation := make(chan string)
	go Start(scanner, confirmation)
	<-confirmation

	// Wait for 4 beats so that user can prep.
	// This is essentially a count-in, and prints
	// asynchronously as to not cause race conditions.
	go fmt.Print("Ready...")
	time.Sleep(bpm.ToDuration(4))
	go fmt.Println("Go!")

	// Context used to automatically close channel and
	// stop scanning after duration.
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	input := make(chan string)   // Input channel used to communicate with reader routine
	go Scan(scanner, input, ctx) // Start scanning

	var inputs []string // Raw string inputs (microsecond timestamps) from Arduino
	// Reads hits until `input` is closed by reader routine
	for num := range input {
		inputs = append(inputs, num)
		go fmt.Print(".")
	}

	fmt.Println()

	cancel()

	// Get parsed output of hits and return it;
	// if there are errors while parsing, it
	// returns a nil array ref and the error.
	return parsers.ParseHits(inputs, bpm, beats)
}

// Scans serial port continuously until context is up.
func Scan(s *bufio.Scanner, hits chan string, ctx context.Context) {
	for s.Scan() {
		select {
		// Closes channel when context is finished
		case <-ctx.Done():
			close(hits)
			return
		// Otherwise block until scanner has text.
		default:
			hits <- s.Text()
		}
	}
	if s.Err() != nil {
		fmt.Println("There was a problem scanning your hits.")
		close(hits)
		runtime.Goexit() // Kill routine if scanner has error
	}
}

// Blocks thread until confirmation is received in the form
// of a hit from the user.
func Start(s *bufio.Scanner, confirmation chan string) {
	// Scan until user hits it
	for s.Scan() {
		confirmation <- s.Text()
		close(confirmation)
		break
	}
	if err := s.Err(); err != nil {
		fmt.Println("Error when scanning for confirmation: ", err.Error())
	}
}
