package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/faxter/pokefx/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokefx > ")
		scanner.Scan()
		userInput := repl.CleanInput(scanner.Text())
		fmt.Println("Your command was:", userInput[0])
	}
}
