package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

var pokedex = map[string]pokeapi.Pokemon{}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the name of the Pomemon to catch.")
	}
	pokemonName := args[0]

	if pokemon, ok := pokedex[pokemonName]; ok {
		fmt.Printf("%s has already been caught!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pokemon, err := pokeapi.GetPokemon(url)
	if err != nil {
		return fmt.Errorf("Error getting Pokemon %s.", pokemonName)
	}

	chance := rand.Float64()
	catchingDifficulity := float64(1) / float64(pokemon.BaseExperience)

	if chance+catchingDifficulity > 0.5 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
