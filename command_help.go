package main

import "fmt"

func commandHelp(_ *config) error {
	output := fmt.Sprint("Welcome to the Pokedex!\nUsage:\n\n")

	for _, v := range commands {
		output += fmt.Sprintf("%s: %s\n", v.name, v.description)
	}

	fmt.Print(output)

	return nil
}
