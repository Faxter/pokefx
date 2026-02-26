package repl

import (
	"fmt"
	"os"
)

func (r *Repl) RegisterCommands() {
	r.registerCommand("exit", "Exit the program", r.commandExit)
	r.registerCommand("help", "Display usage information", r.commandHelp)
}

func (r *Repl) registerCommand(name, description string, callback func() error) {
	r.commands[name] = cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}

func (r *Repl) commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (r *Repl) commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for name, cmd := range r.commands {
		fmt.Printf("\t%s:\t%s\n", name, cmd.description)
	}
	return nil
}

func (r *Repl) ExecuteCommand(command string) {
	cmd, ok := r.commands[command]
	if !ok {
		fmt.Println("unknown command")
	} else {
		err := cmd.callback()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}
