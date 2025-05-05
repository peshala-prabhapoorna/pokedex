package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Error requesting data: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var locationAreas LocationAreas
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf(
			"Error decoding response body: %w", err,
		)
	}

	return locationAreas, nil
}
