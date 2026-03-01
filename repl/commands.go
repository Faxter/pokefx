package repl

import (
	"fmt"
	"os"

	"github.com/faxter/pokefx/internal/pokeapi"
)

const (
	API_MAP_BASE = "https://pokeapi.co/api/v2/location-area/"
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

func (r *Repl) commandMap(cfg *Config, param string) error {
	if param != "" {
		return r.exploreSpecificMap(param)
	}
	var url string
	if cfg.NextPage != "" {
		url = cfg.NextPage
	} else {
		url = API_MAP_BASE + "?offset=0&limit=20"
	}
	response, err := r.retrieveMapListResponse(url)
	if err != nil {
		return err
	}
	for _, name := range response.ExtractMapNames() {
		fmt.Println(name)
	}
	cfg.NextPage = response.Next
	cfg.PreviousPage = response.Previous
	return nil
}

func (r *Repl) exploreSpecificMap(mapname string) error {
	url := API_MAP_BASE + mapname
	response, err := r.retrieveSpecificMapResponse(url)
	if err != nil {
		return err
	}
	for _, pokemon := range response.ExtractPokemonEncounters() {
		fmt.Println(pokemon)
	}
	return nil
}

func (r *Repl) retrieveSpecificMapResponse(url string) (pokeapi.SpecificMapResponse, error) {
	var responseData []byte
	var err error
	cachedValue, foundInCache := r.cache.Get(url)
	if foundInCache {
		responseData = cachedValue
	} else {
		call := pokeapi.CreateApiCall(url)
		responseData, err = call.SendRequest()
		if err != nil {
			return pokeapi.SpecificMapResponse{}, err
		}
		r.cache.Add(url, responseData)
	}

	return pokeapi.ConvertSpecificMapResponseToJson(responseData), nil
}

func (r *Repl) retrieveMapListResponse(url string) (pokeapi.MapListResponse, error) {
	var responseData []byte
	var err error
	cachedValue, foundInCache := r.cache.Get(url)
	if foundInCache {
		responseData = cachedValue
	} else {
		call := pokeapi.CreateApiCall(url)
		responseData, err = call.SendRequest()
		if err != nil {
			return pokeapi.MapListResponse{}, err
		}
		r.cache.Add(url, responseData)
	}

	return pokeapi.ConvertMapListResponseToJson(responseData), nil
}

func (r *Repl) commandMapBack(cfg *Config, _ string) error {
	var url string
	if cfg.PreviousPage != "" {
		url = cfg.PreviousPage
	} else {
		return fmt.Errorf("you're on the first page")
	}
	response, err := r.retrieveMapListResponse(url)
	if err != nil {
		return err
	}
	for _, name := range response.ExtractMapNames() {
		fmt.Println(name)
	}
	cfg.NextPage = response.Next
	cfg.PreviousPage = response.Previous
	return nil
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
