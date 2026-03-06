package repl

import (
	"fmt"
)

func (r *Repl) RegisterCommands() {
	r.registerCommand("exit", "Exit the program", r.commandExit)
	r.registerCommand("help", "Display usage information", r.commandHelp)
	r.registerCommand("map", "List sections of areas, use again to get next page of areas - add map name to get encounters", r.commandMap)
	r.registerCommand("mapb", "List previous page of areas", r.commandMapBack)
	r.registerCommand("explore", "List pokemon that can be encountered in given area", r.commandExplore)
	r.registerCommand("catch", "Try to catch given pokemon - will be added to your pokedex", r.commandCatch)
	r.registerCommand("inspect", "Display info of given pokemon out of your pokedex", r.commandInspect)
	r.registerCommand("pokedex", "Display index of pokemon you have caught", r.commandPokedex)
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
