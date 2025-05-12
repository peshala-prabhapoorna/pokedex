package main

import "github.com/peshala-prabhapoorna/pokedex/internal/pokeapi"

func main() {
	config := config{
		previous: "",
		next:     pokeapi.EndpointLocationArea,
		pokedex:  map[string]pokeapi.Pokemon{},
	}

	startRepl(&config)
}
