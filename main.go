package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	words := strings.Fields(loweredText)
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		formattedInput := cleanInput(input)
		if len(formattedInput) > 0 {
			fmt.Printf("Your command was: %s\n", formattedInput[0])
		} else {
			fmt.Println("Unknow command.")
		}

	}
}
