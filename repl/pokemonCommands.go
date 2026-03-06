package repl

import (
	"fmt"
	"math/rand"

	"github.com/faxter/pokefx/internal/pokeapi"
)

const (
	API_POKEMON_BASE = "https://pokeapi.co/api/v2/pokemon/"
)

func (r *Repl) commandCatch(_ *Config, pokemonName string) error {
	if len(pokemonName) <= 0 {
		return fmt.Errorf("this command needs a pokemon name!")
	}
	if r.pokedex.Check(pokemonName) {
		return fmt.Errorf("you already have %s in your pokedex!", pokemonName)
	}
	url := API_POKEMON_BASE + pokemonName
	data, err := r.fetchRawData(url)
	if err != nil {
		return err
	}
	response, err := decode[pokeapi.PokemonResponse](data)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	cought := tryToCatch(response.BaseExperience)
	if cought {
		fmt.Printf("%s caught! Adding info to pokedex...\n", pokemonName)
		r.pokedex.Add(response)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func tryToCatch(baseValue int) bool {
	// rattata: 51, Lugia: 306
	x := rand.Intn(baseValue)
	return x <= 50
}

func (r *Repl) commandInspect(_ *Config, pokemonName string) error {
	if !r.pokedex.Check(pokemonName) {
		return fmt.Errorf("you have not caught %s yet!", pokemonName)
	}
	poke := r.pokedex.Index[pokemonName]
	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weigth:", poke.Weight)
	fmt.Println("Types:")
	for _, typeEntry := range poke.Types {
		fmt.Println("  -", typeEntry.Type.Name)
	}
	fmt.Println("Moves:")
	for _, move := range poke.Moves {
		fmt.Println("  -", move.Move.Name)
	}
	return nil
}

func (r *Repl) commandPokedex(_ *Config, _ string) error {
	pokedex := r.pokedex.Index
	if len(pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}
	fmt.Println("Your pokedex:")
	for pokeName, _ := range pokedex {
		fmt.Println("  -", pokeName)
	}
	return nil
}
