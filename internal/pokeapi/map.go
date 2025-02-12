package pokeapi

import (
	"net/http"
	"encoding/json"
)

// autogenerated: https://mholt.github.io/json-to-go/
type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string   `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Functions must be capitalized
func GetMap(url string) (LocationArea, error) {
	var result LocationArea
	request, err := http.NewRequest("GET", url, nil)
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