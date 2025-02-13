// Make REPL as a new file
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*Configuration, []string) error
}

type Configuration struct {
	next *string
	previous *string
	client pokeapi.Client
}

func startRepl(configuration Configuration) {

	// Creating input scanner to read from user input
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	// Inf. for loop to poll user input
	for ;; {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			results := cleanInput(scanner.Text())

			if len(results) == 0 {
				fmt.Println("No command entered")
				continue
			}

			command := results[0]
			var parameters []string

			if len(results) > 1 {
				parameters = results[1:]
			}
			
			
			// Checking user input
			if res, ok := commands[command]; ok {
				res.callback(&configuration, parameters)
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
		"map": cliCommand {
			name: "map",
			description: "Lists all the available locations",
			callback: commandMap,
		},
		"mapb": cliCommand {
			name: "mapb",
			description: "Lists all the available locations (previous)",
			callback: commandMapb,
		},
		"explore": cliCommand {
			name: "explore",
			description: "Lists all pokemon in a location area (requires a location area parameter)",
			callback: commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	
	text = strings.ToLower(text)
	splits := strings.Fields(text)

	return splits
}