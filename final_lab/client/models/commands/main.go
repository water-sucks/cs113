package commands

type Command int

type MainCommand Command

const (
	Start MainCommand = iota
	Settings
	Instructions
	Help
	Quit
	InvalidMainCommand
)
