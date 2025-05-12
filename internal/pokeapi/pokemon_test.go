package pokeapi

import "testing"

func TestGetPokemon(t *testing.T) {
	cases := []struct {
		input        string
		expectedName string
	}{
		{
			input:        "https://pokeapi.co/api/v2/pokemon/pikachu",
			expectedName: "pikachu",
		},
		{
			input:        "https://pokeapi.co/api/v2/pokemon/clefairy",
			expectedName: "clefairy",
		},
	}

	for _, c := range cases {
		actual, err := GetPokemon(c.input)

		if err != nil {
			t.Errorf("Error getting Pokemon: %s", c.expectedName)
		}

		if actual.Name != c.expectedName {
			t.Errorf(
				"Fetched Pokemon (%s) is not the expected Pokemon (%s).",
				actual.Name,
				c.expectedName,
			)
		}
	}
}
