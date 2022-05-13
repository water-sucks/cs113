package handlers

import (
	"drum-hero/models"
	"fmt"
)

func MainHelp() {
	fmt.Println(`List of commands:
  start - start a Drum Hero round
  settings - change user settings for rounds
  instructions - show game instructions
  help - show this menu
  quit - quit game and exit`)
}

func SettingsHelp() {
	fmt.Println(`List of settings commands:
  change - change setting
  view - view current settings
  save - save settings and return to main menu
  quit - quit settings and return to main menu
  help - show this menu`)
}

func Instructions() {
	fmt.Println(
		`This is a game that works in a similar way to Guitar Hero.
The goal is to replicate a pattern shown on the screen by
hitting the pad provided at the marked spot in musical time.
If you do it perfectly, you win the round. If not, then you
lose. You can play as many rounds as you like, change settings,
or do otherwise as the game permits. Have fun!`)
}

func StartInstructions(level models.Difficulty) {
	fmt.Printf(
		`When you are ready, confirm with Y, and then hit
the pad to start a count-in of 4 counts. The speeds are as follows
for each difficulty level:
  - Easy: 90 BPM
  - Normal: 120 BPM
  - Hard: 150 BPM
You are currently playing at difficulty level "%s".
Play the entire pattern correctly at this speed in order to win.
If you find find that it hangs after playing the pattern, hit it
one more time after, and it will register that the round ended.
Hint: Using a metronome really helps timing perfectly!

`, level.ToString())
}
