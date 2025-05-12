package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreaPokemons struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetAreaPokemons(url string) (AreaPokemons, error) {
	data, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return AreaPokemons{}, fmt.Errorf("Error requesting data: %w", err)
		}
		defer res.Body.Close()

		dataRes, err := io.ReadAll(res.Body)
		if err != nil {
			return AreaPokemons{}, fmt.Errorf(
				"Error reading response body: %w", err,
			)
		}

		cache.Add(url, dataRes)
		data = dataRes
	}

	var areaPokemons AreaPokemons
	if err := json.Unmarshal(data, &areaPokemons); err != nil {
		return AreaPokemons{}, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return areaPokemons, nil
}
