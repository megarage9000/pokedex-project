// Separating commands into files
package main

import "fmt"

func commandHelp(config *Configuration) error {
	commands := getCommands()
	fmt.Print(
`
Welcome to the Pokedex!
Usage:
`)
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}