package main

import (
	"fmt"
)

func commandPokedex(config *Configuration, params []string) error {
	fmt.Println("Your Pokedex:")
	for name := range config.caughtPokemon {
		fmt.Printf(" -%s\n", name)
	}
	return nil
}