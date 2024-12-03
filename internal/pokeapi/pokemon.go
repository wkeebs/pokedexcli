package pokeapi

import (
	"fmt"
)

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) String() string {
	pokemonStr := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)

	pokemonStr += fmt.Sprintf("Stats:\n")
	for _, s := range p.Stats {
		pokemonStr += fmt.Sprintf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	pokemonStr += fmt.Sprintf("Types:\n")
	for _, t := range p.Types {
		pokemonStr += fmt.Sprintf("  - %s\n", t.Type.Name)
	}

	return pokemonStr
}
