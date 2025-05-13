package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Display Pokemons in a location area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays data of a Pokemon in the Pokedex",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display Pokemons that have been caught",
			callback:    commandPokedex,
		},
	}
}
