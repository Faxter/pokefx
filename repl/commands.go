package repl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/faxter/pokefx/internal/pokeapi"
)

func (r *Repl) RegisterCommands() {
	r.registerCommand("exit", "Exit the program", r.commandExit)
	r.registerCommand("help", "Display usage information", r.commandHelp)
	r.registerCommand("map", "List sections of areas, such as floors in a building or cave - use again to get next page of areas", r.commandMap)
	r.registerCommand("mapb", "List previous page of areas", r.commandMapBack)
}

func (r *Repl) registerCommand(name, description string, callback func(*Config, string) error) {
	r.commands[name] = cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}

func (r *Repl) ExecuteCommand(command string, param string) {
	cmd, ok := r.commands[command]
	if !ok {
		fmt.Println("unknown command")
	} else {
		err := cmd.callback(&r.config, param)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}

func (r *Repl) commandExit(cfg *Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (r *Repl) commandHelp(cfg *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for name, cmd := range r.commands {
		fmt.Printf("\t%s:\t%s\n", name, cmd.description)
	}
	return nil
}

func (r *Repl) fetchRawData(url string) ([]byte, error) {
	if cachedValue, found := r.cache.Get(url); found {
		return cachedValue, nil
	}
	call := pokeapi.CreateApiCall(url)
	responseData, err := call.SendRequest()
	if err != nil {
		return nil, err
	}
	r.cache.Add(url, responseData)
	return responseData, nil
}

func decode[T any](raw []byte) (T, error) {
	var jsonData T
	err := json.Unmarshal(raw, &jsonData)
	return jsonData, err
}
