package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/faxter/pokefx/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cli := repl.CreateRepl()
	cli.RegisterCommands()
	for {
		fmt.Print("pokefx > ")
		scanner.Scan()
		userInput := repl.CleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		parameter := ""
		if len(userInput) > 1 {
			parameter = userInput[1]
		}
		cli.ExecuteCommand(userInput[0], parameter)
	}
}
