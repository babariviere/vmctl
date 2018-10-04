package main

// Command is a command line subcommand
type Command interface {
	Parse([]string) error
	Usage() string
	Spawn() error
}
