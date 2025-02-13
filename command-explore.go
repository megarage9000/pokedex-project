package main

import (
	"fmt"
)

func commandExplore(config *Configuration, params []string) error {

	if len(params) == 0 {
		fmt.Println("No location area indicated to explore!")
		return nil
	}

	fmt.Printf("Exploring %s...\n", params[0])
	result, err := config.client.ExploreLocation(params[0])
	if err != nil {
		fmt.Printf("Error in calling explore: %s", err)
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEnounter := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEnounter.Pokemon.Name)
	}
	return nil
}