package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := Commands()

	for {
		fmt.Printf("Pokedex > ")

		// get input
		key, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		key = strings.TrimSpace(key) // trim newline and spaces

		// get the command
		cmd, ok := commands[key]
		if !ok {
			fmt.Println("Command does not exist! Try 'help'")
			continue
		}

		// execute
		cmd.callback()
	}
}
