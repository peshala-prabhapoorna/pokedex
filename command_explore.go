package main

import (
	"fmt"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the location area to explore.")
	}
	locationArea := args[0]

	url := pokeapi.EndpointLocationArea + "/" + locationArea
	areaPokemons, err := pokeapi.GetAreaPokemons(url)
	if err != nil {
		return fmt.Errorf("Error getting Pokemons in %s.", locationArea)
	}

	var pokemons string
	for _, pokemonEncounter := range areaPokemons.PokemonEncounters {
		pokemons += " - " + pokemonEncounter.Pokemon.Name + "\n"
	}

	fmt.Print(
		"Exploring "+locationArea+"...\n",
		"Found Pokemon:\n",
		pokemons,
	)

	return nil
}
