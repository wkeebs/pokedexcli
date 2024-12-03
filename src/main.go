package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/wkeebs/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient     pokeapi.Client
	locationPageIndex int
	exploreAreaName   string
}

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient:     client,
		locationPageIndex: -1,
	}
	reader := bufio.NewReader(os.Stdin)
	commands := Commands(cfg)

	for {
		fmt.Printf("Pokedex > ")

		// get input
		userInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		userInput = strings.TrimSpace(userInput) // trim newline and spaces
		userInputSplit := FilterEmpty(strings.Split(userInput, " "))

		args := userInputSplit[1:]

		// get the command
		cmd, ok := commands[userInputSplit[0]]
		if !ok {
			fmt.Println("Command does not exist! Try 'help'")
			continue
		}

		// cfg.pokeapiClient.Cache.PrintCache()

		// execute
		err = cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}
