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

		command := words[0]
		isCommand := false
		for c, v := range Commands {
			if c == command {
				isCommand = true

				if err := v.callback(); err != nil {
					fmt.Println(err)
				}
			}
		}

		if !isCommand {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	textLowerCase := strings.ToLower(text)
	output := strings.Fields(textLowerCase)
	return output
}
