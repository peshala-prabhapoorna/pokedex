package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
	config      *config
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
			config:      nil,
		},
		"explore": {
			name:        "explore",
			description: "Display Pokemons in a location area",
			callback:    commandExplore,
			config:      nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
			config:      &configMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
			config:      &configMap,
		},
	}
}
