package repl

import (
	"strings"
	"time"

	"github.com/faxter/pokefx/internal/pokecache"
)

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
		cache:    *pokecache.NewCache(5 * time.Second),
	}
}

type Repl struct {
	commands map[string]cliCommand
	config   Config
	cache    pokecache.Cache
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
