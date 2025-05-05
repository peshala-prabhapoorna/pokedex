package main

type config struct {
	Previous string
	Next     string
}

var configMap = config{
	Previous: "",
	Next:     "https://pokeapi.co/api/v2/location-area",
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
