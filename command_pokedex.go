package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Command doesn't accept arguments.")
	}

	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemons.")
		return nil
	}

	pokemons := ""
	for pokemonName := range cfg.pokedex {
		pokemons += fmt.Sprintf("  - %s\n", pokemonName)
	}

	fmt.Print("Your Pokedex:\n", pokemons)

	return nil
}
