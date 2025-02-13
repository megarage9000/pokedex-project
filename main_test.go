package main

import (
	"testing"
)

// - All tests must be prefixed with "Test"
// - All test must have "t *testing.T" as a parameter
func TestCleanInput(t *testing.T) {

	// Creating test cases, here we used anonymous struct
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "lego rocks",
			expected: []string{"lego", "rocks"},
		},
	}

	// Run test cases
	for _, c := range cases {

		actual := cleanInput(c.input)
		
		// Ensuring length is correct
		if len(actual) != len(c.expected) {
			t.Errorf(`
			-----------------
			TEST FAILED:
			expected length: %v
			actual length: %v
		`, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			// Ensuring words are correct
			if expected != word {
				t.Errorf(`
					-----------------
					TEST FAILED:
					expected: %v
					actual: %v
				`, expected, word)
			}
		}
	}
}

