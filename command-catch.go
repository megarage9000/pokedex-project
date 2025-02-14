package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandCatch(config *Configuration, params []string) error {

	if len(params) == 0 {
		fmt.Println("No pokemon listed to catch!")
		return nil
	}

	pokemon, err := config.client.GetPokemonInformation(params[0])
	if err != nil {
		fmt.Printf("Unable to get information on pokemon %s!\n", params[0])
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if caught := pokeapi.CatchPokemon(pokemon); !caught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}