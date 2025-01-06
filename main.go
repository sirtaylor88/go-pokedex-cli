package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	words := strings.Fields(loweredText)
	return words
}

func main() {
	fmt.Println("Hello, World!")
}
