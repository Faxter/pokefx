package repl

import (
	"strings"
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
	}
}

type Repl struct {
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
