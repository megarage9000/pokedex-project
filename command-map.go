package main

import (
	"internal/pokeapi"
	"internal/pokecache"
	"fmt"
	"encoding/json"
)

func commandMap(config *Configuration) error {
	
	defaultNext := "https://pokeapi.co/api/v2/location-area"

	// check if the next url is set
	if config.next == nil {
		config.next = &defaultNext
	}

	command := *config.next

	locationArea, ok := getCachedLocationAreaResult(&config.cache, command)
	if !ok {
		response, err := pokeapi.ListLocations(command)
		if err != nil {
			return err
		} else {
			locationArea = response
		}
	} 

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	// Cache data
	cachedData, err := json.Marshal(locationArea)
	if err == nil {
		config.cache.Add(command, cachedData)
	}

	listLocationAreas(locationArea.Results)

	return nil
}

func commandMapb(config *Configuration) error {

	// check if the previous url is set
	if config.previous == nil{
		return nil
	}

	command := *config.previous

	locationArea, ok := getCachedLocationAreaResult(&config.cache, command)
	if !ok {
		response, err := pokeapi.ListLocations(command)
		if err != nil {
			return err
		} else {
			locationArea = response
		}
	} 

	config.next = locationArea.Next
	config.previous = locationArea.Previous

	// Cache data
	cachedData, err := json.Marshal(locationArea)
	if err == nil {
		config.cache.Add(command, cachedData)
	}

	listLocationAreas(locationArea.Results)

	return nil
}

func listLocationAreas(locationAreaResults []pokeapi.LocationAreaResult) {
	for _, locationArea := range locationAreaResults{
		fmt.Printf("%s\n", locationArea.Name)
	}
}

func getCachedLocationAreaResult(cache *pokecache.Cache, command string) (pokeapi.LocationArea, bool) {
	result, ok := cache.Get(command)
	if !ok {
		fmt.Printf("CACHE unable to get data for %s\n", command)
		return pokeapi.LocationArea{}, false
	}
	var locationArea pokeapi.LocationArea

	if err := json.Unmarshal(result, &locationArea); err != nil {
		fmt.Println("CACHE unable to unmarshal data")
		return pokeapi.LocationArea{}, false
	}
	fmt.Println("CACHE got cached result!")
	return locationArea, true
}