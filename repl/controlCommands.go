package repl

import (
	"fmt"
	"os"
)

func (r *Repl) commandExit(_ *Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (r *Repl) commandHelp(_ *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	maxCommandLength := 0
	for command, _ := range r.commands {
		if len(command) > maxCommandLength {
			maxCommandLength = len(command)
		}
	}
	for name, cmd := range r.commands {
		fmt.Printf("\t%-*s%s\n", maxCommandLength+2, name, cmd.description)
	}
	return nil
}
