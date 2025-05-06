package pokeapi

import (
	"time"

	"github.com/peshala-prabhapoorna/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(8 * time.Second)
