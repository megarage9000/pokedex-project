package pokeapi

import (
	"net/http"
	"encoding/json"
)

// Functions must be capitalized
func (c * Client) ListLocations(locationUrl string) (LocationArea, error) {

	getUrl := baseUrl + "/location-area"
	if locationUrl != "" {
		getUrl = locationUrl
	}

	var result LocationArea
	
	// Get Cache
	if cacheResult, ok := c.cache.Get(getUrl); ok {
		if err := json.Unmarshal(cacheResult, &result); err != nil {
			return result, err
		}
		return result, nil
	} 
	
	// Perform http get on location area
	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return result, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	// Cache data
	if cachedData, err := json.Marshal(response); err == nil {
		c.cache.Add(getUrl, cachedData)
	}

	return result, err
}