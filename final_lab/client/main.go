package main

import (
	"drum-hero/cli"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Executes terminal command `toilet`, which may not
	// be available on your system; this is for vanity's
	// sake and displaying the title in large letters.
	out, _ := exec.Command("toilet", "Drum Hero").Output()
	fmt.Printf("%s", out)

	fmt.Printf("Welcome to Drum Hero!\n\n")

	// Shell is an infinite loop, this is idiomatic
	err := cli.Shell()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s", err.Error())
		os.Exit(1)
	}
}
