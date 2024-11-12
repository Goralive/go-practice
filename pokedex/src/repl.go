package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
}

func startRepl() {
	scaner := bufio.NewScanner(os.Stdin)
	for {
		if scaner.Scan() != false {
			userInput := scaner.Text()
			fmt.Printf("Pokedex -> %s\n", userInput)
			menuItem, exists := getCommands()[userInput]
			if exists {
				if err := menuItem.callback(); err != nil {
					fmt.Errorf("Something go wrong %v\n", err)
				}
			}
		}
	}
}
