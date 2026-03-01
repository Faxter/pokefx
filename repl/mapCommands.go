package repl

import (
	"fmt"

	"github.com/faxter/pokefx/internal/pokeapi"
)

const (
	API_MAP_BASE = "https://pokeapi.co/api/v2/location-area/"
)

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
	data, err := r.fetchRawData(url)
	if err != nil {
		return err
	}
	response, err := decode[pokeapi.MapListResponse](data)
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
	fmt.Printf("Exploring %s...\n", mapname)
	url := API_MAP_BASE + mapname
	data, err := r.fetchRawData(url)
	if err != nil {
		return err
	}
	response, err := decode[pokeapi.SpecificMapResponse](data)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.ExtractPokemonEncounters() {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}

func (r *Repl) commandMapBack(cfg *Config, _ string) error {
	var url string
	if cfg.PreviousPage != "" {
		url = cfg.PreviousPage
	} else {
		return fmt.Errorf("you're on the first page")
	}
	data, err := r.fetchRawData(url)
	if err != nil {
		return err
	}
	response, err := decode[pokeapi.MapListResponse](data)
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
