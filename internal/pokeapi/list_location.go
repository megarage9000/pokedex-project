package pokeapi

import (
	"net/http"
	"encoding/json"
)

// Functions must be capitalized
func ListLocations(url string) (LocationArea, error) {

	getUrl := baseUrl + "/location-area"
	if url != "" {
		getUrl = url
	}

	var result LocationArea
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

	return result, err
}