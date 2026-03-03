package pokedex

import "github.com/faxter/pokefx/internal/pokeapi"

type Pokedex struct {
	Index map[string]pokeapi.PokemonResponse
}

func CreatePokedex() Pokedex {
	return Pokedex{Index: make(map[string]pokeapi.PokemonResponse)}
}

func (p *Pokedex) Add(pokemon pokeapi.PokemonResponse) {
	p.Index[pokemon.Name] = pokemon
}

func (p *Pokedex) Check(name string) bool {
	_, found := p.Index[name]
	return found
}
