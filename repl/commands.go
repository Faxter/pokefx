package repl

import (
	"fmt"
	"os"

	"github.com/faxter/pokefx/pokeapi"
)

func (r *Repl) RegisterCommands() {
	r.registerCommand("exit", "Exit the program", r.commandExit)
	r.registerCommand("help", "Display usage information", r.commandHelp)
	r.registerCommand("map", "List sections of areas, such as floors in a building or cave - use again to get next page of areas", r.commandMap)
	r.registerCommand("mapb", "List previous page of areas", r.commandMapBack)
}

func (r *Repl) registerCommand(name, description string, callback func(*Config) error) {
	r.commands[name] = cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}

func (r *Repl) commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (r *Repl) commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for name, cmd := range r.commands {
		fmt.Printf("\t%s:\t%s\n", name, cmd.description)
	}
	return nil
}

func (r *Repl) commandMap(cfg *Config) error {
	var call pokeapi.ApiCall
	if cfg.NextPage != "" {
		call = pokeapi.CreateApiCall(cfg.NextPage)
	} else {
		call = pokeapi.CreateApiCall("https://pokeapi.co/api/v2/location-area/")
	}
	response, err := call.SendRequest()
	if err != nil {
		return err
	}
	for _, name := range response.ExtractNames() {
		fmt.Println(name)
	}
	cfg.NextPage = response.Next
	cfg.PreviousPage = response.Previous
	return nil
}

func (r *Repl) commandMapBack(cfg *Config) error {
	var call pokeapi.ApiCall
	if cfg.PreviousPage != "" {
		call = pokeapi.CreateApiCall(cfg.PreviousPage)
	} else {
		return fmt.Errorf("you're on the first page")
	}
	response, err := call.SendRequest()
	if err != nil {
		return err
	}
	for _, name := range response.ExtractNames() {
		fmt.Println(name)
	}
	cfg.NextPage = response.Next
	cfg.PreviousPage = response.Previous
	return nil
}

func (r *Repl) ExecuteCommand(command string) {
	cmd, ok := r.commands[command]
	if !ok {
		fmt.Println("unknown command")
	} else {
		err := cmd.callback(&r.config)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}
