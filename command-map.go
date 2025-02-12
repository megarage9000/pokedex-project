package main

import (
	"internal/pokeapi"
	"fmt"
)

func commandMap(config *Configuration) error {
	
	defaultNext := "https://pokeapi.co/api/v2/location/"

	// check if the next url is set
	if config.next == nil {
		config.next = &defaultNext
	}

	locationAreaResults, err := pokeapi.GetMap(*config.next)
	if err != nil {
		return err
	}

	config.next = locationAreaResults.Next
	config.previous = locationAreaResults.Previous

	for _, locationArea := range locationAreaResults.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}



func commandMapb(config *Configuration) error {

	// check if the previous url is set
	if config.previous == nil{
		return nil
	}
	return nil
}