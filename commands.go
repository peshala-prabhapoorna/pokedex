package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Displays location areas",
			callback:    commandMap,
			config:      &configMap,
		},
	}
}
