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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var table = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		formattedInput := cleanInput(input)
		if len(formattedInput) > 0 {
			cmd := formattedInput[0]
			command, ok := table[cmd]
			if !ok {
				fmt.Println("Unknow command.")
			}
			command.callback()
			if cmd == "help" {
				fmt.Println("")
				for key, value := range table {
					fmt.Printf("%s: %s\n", key, value.description)
				}
			}
		} else {
			fmt.Println("Unknow command.")
		}

	}
}
