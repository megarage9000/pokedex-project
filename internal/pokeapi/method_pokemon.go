package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
	"io"
	"math/rand"
)

func (c *Client) GetPokemonInformation(pokemonName string) (Pokemon, error) {

	var pokemon Pokemon

	if pokemonName == "" {
		return pokemon, errors.New("no pokemon name provided")
	}

	getUrl := baseUrl + "/pokemon/" + pokemonName

	// Get from cache
	if cacheResult, ok := c.cache.Get(getUrl); ok {
		if err := json.Unmarshal(cacheResult, &pokemon); err != nil {
			return pokemon, fmt.Errorf("error in getting from cache: %s", err)
		}
		return pokemon, nil
	}


	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return pokemon, fmt.Errorf("error in making request for %s: %s", getUrl, err)
	}


	response, err := c.httpClient.Do(request)
	if err != nil {
		return pokemon, fmt.Errorf("error in performing GET request to %s: %s", getUrl, err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return pokemon, fmt.Errorf("error in reading GET response from %s: %s", getUrl, err)
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, fmt.Errorf("error in decoding data from response %s: %s", getUrl, err)
	}

	c.cache.Add(getUrl, data)

	return pokemon, nil
}


// For now all pokemon have 50% catch rate
func CatchPokemon(pokemon Pokemon) bool {
	baseExperience := pokemon.BaseExperience
	n := rand.Intn(baseExperience)
	return n <= int(baseExperience / 2)
}