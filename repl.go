package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	words := strings.Fields(lowerCaseText)
	return words
}
