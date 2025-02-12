// Separating commands into files
package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println(`
Welcome to the Pokedex!
Usage:
`)
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}