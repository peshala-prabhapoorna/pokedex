package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the name of the Pomemon to catch.")
	}
	pokemonName := args[0]

	if pokemon, ok := cfg.pokedex[pokemonName]; ok {
		fmt.Printf("%s has already been caught!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := pokeapi.EndpointPokemon(pokemonName)
	pokemon, err := pokeapi.GetPokemon(url)
	if err != nil {
		return fmt.Errorf("Error getting Pokemon %s.", pokemonName)
	}

	chance := rand.Float64()
	catchingDifficulity := float64(1) / float64(pokemon.BaseExperience)

	if chance+catchingDifficulity > 0.5 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
