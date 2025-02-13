package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)

func (c *Client) ExploreLocation(location string) (LocationPokemonList, error) {

	getUrl := baseUrl + "/location-area/" + location

	var result LocationPokemonList

	// Get from cache
	if cacheResult, ok := c.cache.Get(getUrl); ok {
		if err := json.Unmarshal(cacheResult, &result); err != nil {
			return result, fmt.Errorf("error in getting from cache: %s", err)
		}
		return result, nil
	}

	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return result, fmt.Errorf("error in getting from making GET request to %s: %s", getUrl, err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return result, fmt.Errorf("error in performing GET request to %s: %s", getUrl, err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return result, fmt.Errorf("error in reading GET response from %s: %s", getUrl, err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("error in decoding data from response %s: %s", getUrl, err)
	}

	// Save to cache
	c.cache.Add(getUrl, data)

	return result, nil
} 