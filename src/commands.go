package main

import (
	"fmt"
	"os"
)

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

func commandMap() error {
	areas, err := getAreas("https://pokeapi.co/api/v2/location-area/", 20, 0)
	if err != nil {
		return err
	}

	for _, a := range areas {
		fmt.Println(a.Name)
	}
	return nil
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapB,
		},
	}
}
