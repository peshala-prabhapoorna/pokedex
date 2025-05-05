package main

type config struct {
	Previous string
	Next     string
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
