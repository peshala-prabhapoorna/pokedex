package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var parameter string
		if len(words) == 2 {
			parameter = words[1]
		}

		command, isCommand := commands[commandName]
		if isCommand {
			if err := command.callback(command.config, parameter); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	textLowerCase := strings.ToLower(text)
	output := strings.Fields(textLowerCase)
	return output
}
