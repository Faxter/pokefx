package repl

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/faxter/pokefx/internal/pokeapi"
	"github.com/faxter/pokefx/internal/pokecache"
	"github.com/faxter/pokefx/internal/pokedex"
)

type Repl struct {
	commands map[string]cliCommand
	config   Config
	cache    pokecache.Cache
	pokedex  pokedex.Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

type Config struct {
	NextPage     string
	PreviousPage string
}

func CleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}

func CreateRepl() *Repl {
	return &Repl{
		commands: make(map[string]cliCommand),
		config:   Config{NextPage: "", PreviousPage: ""},
		cache:    *pokecache.NewCache(60 * time.Second),
		pokedex:  pokedex.CreatePokedex(),
	}
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
