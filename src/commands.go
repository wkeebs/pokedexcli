package main

import (
	"fmt"
	"os"
)

const AREA_LIMIT = 20

type command struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	helpMessage := "Welcome to the Pokedex!\n\n"

	for _, command := range Commands() {
		helpMessage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	fmt.Println(helpMessage)

	return nil
}

func commandExit() error {
	fmt.Println("Powering down...")
	os.Exit(0)
	return nil
}

func commandMap(increment int) func() error {
	areaIndex := 0 // essentially, current page in the API
	return func() error {
		areas, err := getAreas("https://pokeapi.co/api/v2/location-area/", AREA_LIMIT, areaIndex)
		if err != nil {
			return err
		}

		// increment or decrement current index after each call
		// -> changes whether map or mapb is called
		areaIndex += increment

		for _, a := range areas {
			fmt.Println(a.Name)
		}
		return nil
	}
}

func commandMapB() error {
	return nil
}

func Commands() map[string]command {
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
			callback:    commandMap(1),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMap(-1),
		},
	}
}
