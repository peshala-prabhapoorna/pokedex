package pokeapi

import (
	"time"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(60 * time.Second)

var EndpointLocationArea = "https://pokeapi.co/api/v2/location-area"
