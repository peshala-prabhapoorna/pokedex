package main

import (
	"fmt"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the Pokemon to inspect.")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	printPokemonData(pokemon)

	return nil
}

func printPokemonData(pokemon pokeapi.Pokemon) {
	stats := ""
	for _, stat := range pokemon.Stats {
		stats += fmt.Sprintf("\n  -%s: %d", stat.Stat.Name, stat.BaseStat)
	}

	types := ""
	for _, t := range pokemon.Types {
		types += fmt.Sprintf("\n  -%s", t.Type.Name)
	}

	fmt.Print(
		"Name: ", pokemon.Name, "\n",
		"Height: ", pokemon.Height, "\n",
		"Weight: ", pokemon.Weight, "\n",
		"Stats:", stats, "\n",
		"Types:", types, "\n",
	)
}
