package main

import (
	"fmt"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandExplore(_ *config, locationArea string) error {
	url := pokeapi.EndpointLocationArea + "/" + locationArea
	areaPokemons, err := pokeapi.GetAreaPokemons(url)
	if err != nil {
		return fmt.Errorf("Error getting Pokemons: %w", err)
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
