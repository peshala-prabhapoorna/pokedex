package main

import "strings"

func cleanInput(text string) []string {
	textLowerCase := strings.ToLower(text)
	output := strings.Fields(textLowerCase)
	return output
}
