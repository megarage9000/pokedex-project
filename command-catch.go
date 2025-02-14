package main

import (
	"fmt"
)

func commandExplore(config *Configuration, params []string) error {

	if len(params) == 0 {
		fmt.Println("No pokemon listed to catch!")
		return nil
	}

	pokemon, err := config.client.GetPokemonInformation(params[0])
	if err != nil {

	}
	
}