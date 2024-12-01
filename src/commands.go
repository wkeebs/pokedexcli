package main

import (
	"fmt"
	"os"
)

const AREA_LIMIT = 20

type command struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandHelp(cfg *config) error {
	helpMessage := "Welcome to the Pokedex!\n\n"

	for _, command := range Commands(cfg) {
		helpMessage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	fmt.Println(helpMessage)

	return nil
}

func commandExit(cfg *config) error {
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

func commandMapf(cfg *config) error {
	return commandMap(cfg, true)
}

func commandMapb(cfg *config) error {
	return commandMap(cfg, false)
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
	}
}
