package main

import (
	"fmt"
)

func commandInspect(config *Configuration, params []string) error {

	if len(params) == 0 {
		fmt.Println("No pokemon listed to inspect!")
		return nil
	}

	pokemon, ok := config.caughtPokemon[params[0]]
	if !ok {
		fmt.Printf("Pokemon %s has yet to be caught!\n", params[0])
		return nil
	}

	pokemon.InspectPokemon()
	return nil
}