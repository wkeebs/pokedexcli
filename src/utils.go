package main

import (
	"sort"
)

func FilterEmpty(s []string) []string {
	var newSlice []string
	for _, s := range s {
		if s != "" {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func GetSortedCommands(commands map[string]command) []command {
	var cmdSlice []command
	for _, val := range commands {
		cmdSlice = append(cmdSlice, val)
	}
	sort.Slice(cmdSlice, func(i, j int) bool {
		return cmdSlice[i].id < cmdSlice[j].id
	})
	return cmdSlice
}
