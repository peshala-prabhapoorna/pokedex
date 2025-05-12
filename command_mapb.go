package main

import (
	"fmt"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locationAreas, err := pokeapi.GetLocationAreas(cfg.previous)
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w", err)
	}

	updateConfig(cfg, locationAreas.Next, locationAreas.Previous)

	var output string
	for _, locationArea := range locationAreas.Results {
		output += fmt.Sprintln(locationArea.Name)
	}

	fmt.Print(output)

	return nil
}
