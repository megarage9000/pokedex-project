package main

import (
	"internal/pokeapi"
	"fmt"
)

func commandMap(config *Configuration) error {

	var command string

	// check if the next url is set
	if config.next != nil {
		command = *config.next
	}

	locationArea, err := config.client.ListLocations(command)
	if err != nil {
		return err
	}

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	listLocationAreas(locationArea.Results)

	return nil
}

func commandMapb(config *Configuration) error {

	// check if the previous url is set
	if config.previous == nil{
		return nil
	}

	command := *config.previous

	locationArea, err := config.client.ListLocations(command)
	if err != nil {
		return err
	}

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	listLocationAreas(locationArea.Results)

	return nil
}

func listLocationAreas(locationAreaResults []pokeapi.LocationAreaResult) {
	for _, locationArea := range locationAreaResults{
		fmt.Printf("%s\n", locationArea.Name)
	}
}
