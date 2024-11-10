package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	menu()
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func menu() {
	var mainMenu = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	for true {
		scaner := bufio.NewScanner(os.Stdin)
		if scaner.Scan() != false {
			userInput := scaner.Text()
			fmt.Printf("Pokedex -> %s\n", userInput)
			menuItem, exists := mainMenu[userInput]
			if exists {
				if err := menuItem.callback(); err != nil {
					fmt.Errorf("Something go wrong %v\n", err)
				}
			}
		}
	}
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!")
	fmt.Printf("\nUsage:\n")
	fmt.Printf("\n%v: %v", "help", "Displays a help message")
	fmt.Printf("\n%v: %v", "exit", "Exit the Pokedex\n")

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
