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

		// cfg.pokeapiClient.Cache.PrintCache()

		// execute
		err = cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}
