package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
	data, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("Error requesting data: %w", err)
		}
		defer res.Body.Close()

		dataRes, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreas{}, fmt.Errorf(
				"Error reading response body: %w", err,
			)
		}

		cache.Add(url, dataRes)
		data = dataRes
	}

	var locationAreas LocationAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf(
			"Error decoding response body: %w", err,
		)
	}

	return locationAreas, nil
}
