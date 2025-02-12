// Make REPL as a new file
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}


func startRepl() {
	
	// Creating input scanner to read from user input
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	// Inf. for loop to poll user input
	for ;; {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			results := cleanInput(scanner.Text())
			command := results[0]
			
			// Checking user input
			if res, ok := commands[command]; ok {
				res.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

// Dynamically generate commands as a function for all methods to use
func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": cliCommand {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": cliCommand {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	
	text = strings.ToLower(text)
	splits := strings.Fields(text)

	return splits
}