package main

import "github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"

type config struct {
	previous string
	next     string
	pokedex  map[string]pokeapi.Pokemon
}

func updateConfig(config *config, next, previous *string) {
	if next != nil {
		config.next = *next
	} else {
		config.next = ""
	}

	if previous != nil {
		config.previous = *previous
	} else {
		config.previous = ""
	}
}
