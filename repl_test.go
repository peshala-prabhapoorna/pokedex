package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello wORld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO WORLD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("actual input length (%d) is not equal to expected input length (%d)", len(actual), len(c.expected))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("actual input (%s) is not equal to expected input (%s)", actual[i], c.expected[i])
			}
		}
	}
}
