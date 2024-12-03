package main

import (
	"fmt"
	"os"

	"math/rand"

	"github.com/wkeebs/pokedexcli/internal/pokeapi"
)

const AREA_LIMIT = 20

type command struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func commandHelp(cfg *config, _ ...string) error {
	helpMessage := "Welcome to the Pokedex!\n\n"

	for _, command := range Commands(cfg) {
		helpMessage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	fmt.Println(helpMessage)

	return nil
}

func commandExit(cfg *config, _ ...string) error {
	fmt.Println("Powering down...")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, forward bool) error {
	if cfg.locationPageIndex <= 0 && !forward {
		return fmt.Errorf("Already on page 0!")
	}

	var pageIndex int
	if forward {
		pageIndex = cfg.locationPageIndex + 1
	} else {
		pageIndex = cfg.locationPageIndex - 1
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(AREA_LIMIT, pageIndex)
	if err != nil {
		return err
	}

	if forward {
		cfg.locationPageIndex++
	} else {
		cfg.locationPageIndex--
	}

	fmt.Printf("=== LOCATION AREAS - PAGE %d ===\n", pageIndex)
	for _, a := range locationResp.Results {
		fmt.Println(a.Name)
	}
	return nil
}

func commandMapf(cfg *config, _ ...string) error {
	return commandMap(cfg, true)
}

func commandMapb(cfg *config, _ ...string) error {
	return commandMap(cfg, false)
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location to explore! i.e., 'explore pastoria-city-area'")
	}

	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)

	locationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	if len(locationResp.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found!")
	} else {
		fmt.Println("Found Pokemon:")
		for _, p := range locationResp.PokemonEncounters {
			pokemon := p.Pokemon
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon to catch! i.e., 'catch bulbasaur'")
	}

	pokemonName := args[0]

	pokemonResponse, err := cfg.pokeapiClient.GetPokemonSpecies(pokemonName)
	if err != nil {
		return err
	}

	pokemon := pokeapi.Pokemon{
		Name:   pokemonResponse.Name,
		Height: pokemonResponse.Height,
		Weight: pokemonResponse.Weight,
		Stats:  pokemonResponse.Stats,
		Types:  pokemonResponse.Types,
	}

	res := rand.Intn(pokemonResponse.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon to inspect! i.e., 'inspect bulbasaur'")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon '%s' has not been caught!", pokemonName)
	}

	fmt.Print(pokemon.String())

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}

func Commands(cfg *config) map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"quit": {
			name:        "quit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon in a given area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon's stats.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokemon you have caught.",
			callback:    commandPokedex,
		},
	}
}
