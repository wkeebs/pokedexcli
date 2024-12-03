package pokeapi

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Name           string
	BaseExperience int
}

func (p Pokemon) String() string {
	return fmt.Sprintf("%s\n -- Capture Rate: %d", strings.ToTitle(p.Name), p.BaseExperience)
}
