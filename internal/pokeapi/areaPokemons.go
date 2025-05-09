package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreaPokemons struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionEncounterMethod struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod          `json:"encounter_method"`
	VersionDetails  []VersionEncounterMethod `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}
type VersionEncounterDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}
type PokemonEncounters struct {
	Pokemon        Pokemon                   `json:"pokemon"`
	VersionDetails []VersionEncounterDetails `json:"version_details"`
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
