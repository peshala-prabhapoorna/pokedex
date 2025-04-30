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

		fmt.Printf("Your command was: %s\n", words[0])
	}
}

func cleanInput(text string) []string {
	textLowerCase := strings.ToLower(text)
	output := strings.Fields(textLowerCase)
	return output
}
