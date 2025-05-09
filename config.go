package main

import "github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"

type config struct {
	Previous string
	Next     string
}

var configMap = config{
	Previous: "",
	Next:     pokeapi.EndpointLocationArea,
}

func updateConfig(config *config, next, previous *string) {
	if next != nil {
		config.Next = *next
	} else {
		config.Next = ""
	}

	if previous != nil {
		config.Previous = *previous
	} else {
		config.Previous = ""
	}
}
