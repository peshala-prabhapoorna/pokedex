package pokeapi

import (
	"time"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(60 * time.Second)

var endpointLocationArea = "https://pokeapi.co/api/v2/location-area/"
var endpointPokemon = "https://pokeapi.co/api/v2/pokemon/"

func EndpointLocationArea(locationArea *string) string {
	if locationArea != nil {
		return endpointLocationArea + *locationArea
	}

	return endpointLocationArea
}

func EndpointPokemon(pokemonName string) string {
	return endpointPokemon + pokemonName
}
