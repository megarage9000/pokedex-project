package pokeapi

import (
	"net/http"
	"json"
	"io"
)

func (c *Client) ExploreLocation(location string) (LocationPokemonList, error) {

	getUrl := baseUrl + "/location-area/" + location

	var result LocationPokemonList

	// Get from cache
	if cacheResult, ok := c.cache.Get(getUrl); ok {
		if err := json.Unmarshal(cacheResult, &result); err != nil {
			return result, nil
		}
		return result, nil
	}

	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return result, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return result, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	// Save to cache
	c.cache.Add(getUrl, data)

	return result, nil
} 