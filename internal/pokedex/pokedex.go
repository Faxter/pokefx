package pokedex

import "github.com/faxter/pokefx/internal/pokeapi"

type Pokedex struct {
	index map[string]pokeapi.PokemonResponse
}

func CreatePokedex() Pokedex {
	return Pokedex{index: make(map[string]pokeapi.PokemonResponse)}
}

func (p *Pokedex) Add(pokemon pokeapi.PokemonResponse) {
	p.index[pokemon.Name] = pokemon
}

func (p *Pokedex) Check(name string) bool {
	_, found := p.index[name]
	return found
}
