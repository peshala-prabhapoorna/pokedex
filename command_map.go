package main

import (
	"fmt"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandMap(config *config) error {
	locationAreas, err := pokeapi.GetLocationAreas(config.Next)
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w", err)
	}

	updateConfig(config, locationAreas.Next, locationAreas.Previous)

	var output string
	for _, locationArea := range locationAreas.Results {
		output += fmt.Sprintln(locationArea.Name)
	}

	fmt.Print(output)

	return nil
}
