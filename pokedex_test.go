package main

import (
	"testing"
	"internal/pokeapi"
	"time"
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

func TestPokeApiMap(t *testing.T) {

	inputStr := "https://pokeapi.co/api/v2/location-area"
	nextInputStr := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

	cases := []struct {
		input string
		expected pokeapi.LocationArea
	} {
		{
			input: inputStr,
			expected: pokeapi.LocationArea {
				Count: 1089,
				Next: &nextInputStr,
				Previous: nil,
				Results: []pokeapi.LocationAreaResult {
					{
						Name: "canalave-city-area",
						URL: "https://pokeapi.co/api/v2/location-area/1/",
					},
					{
						Name: "eterna-city-area",
						URL: "https://pokeapi.co/api/v2/location-area/2/",
					},
					{
						Name: "pastoria-city-area",
						URL: "https://pokeapi.co/api/v2/location-area/3/",
					},					
					{
						Name: "sunyshore-city-area",
						URL: "https://pokeapi.co/api/v2/location-area/4/",
					},
					{
						Name: "sinnoh-pokemon-league-area",
						URL: "https://pokeapi.co/api/v2/location-area/5/",
					},
					{
						Name: "oreburgh-mine-1f",
						URL: "https://pokeapi.co/api/v2/location-area/6/",
					},
					{
						Name: "oreburgh-mine-b1f",
						URL: "https://pokeapi.co/api/v2/location-area/7/",
					},
					{
						Name: "valley-windworks-area",
						URL: "https://pokeapi.co/api/v2/location-area/8/",
					},
					{
						Name: "eterna-forest-area",
						URL: "https://pokeapi.co/api/v2/location-area/9/",
					},
					{
						Name: "fuego-ironworks-area",
						URL: "https://pokeapi.co/api/v2/location-area/10/",
					},
					{
						Name: "mt-coronet-1f-route-207",
						URL: "https://pokeapi.co/api/v2/location-area/11/",
					},
					{
						Name: "mt-coronet-2f",
						URL: "https://pokeapi.co/api/v2/location-area/12/",
					},
					{
						Name: "mt-coronet-3f",
						URL: "https://pokeapi.co/api/v2/location-area/13/",
					},
					{
						Name: "mt-coronet-exterior-snowfall",
						URL: "https://pokeapi.co/api/v2/location-area/14/",
					},
					{
						Name: "mt-coronet-exterior-blizzard",
						URL: "https://pokeapi.co/api/v2/location-area/15/",
					},
					{
						Name: "mt-coronet-4f",
						URL: "https://pokeapi.co/api/v2/location-area/16/",
					},
					{
						Name: "mt-coronet-4f-small-room",
						URL: "https://pokeapi.co/api/v2/location-area/17/",
					},
					{
						Name: "mt-coronet-5f",
						URL: "https://pokeapi.co/api/v2/location-area/18/",
					},
					{
						Name: "mt-coronet-6f",
						URL: "https://pokeapi.co/api/v2/location-area/19/",
					},
					{
						Name: "mt-coronet-1f-from-exterior",
						URL: "https://pokeapi.co/api/v2/location-area/20/",
					},
				},
			},
		},
	}
	client := pokeapi.NewClient(5 * time.Second, 1 * time.Minute)
	// Run test cases
	for _, testCase := range cases {

		actual, err := client.ListLocations(testCase.input)
		expected := testCase.expected

		if err != nil {
			t.Errorf(
`
-------- TEST FAILED --------
error when calling getMap():
`)
		}

		if actual.Count != expected.Count {
			t.Errorf(
`
-------- TEST FAILED --------
expected Count: %d
actual Count: %d
`,
			expected.Count,
			actual.Count)
		}
		// TODO check if nil
		if (actual.Next != expected.Next && (actual.Next == nil  || expected.Next == nil)){
			t.Errorf(
`
-------- TEST FAILED --------
expected Next: %v
actual Next: %v
`,
			expected.Next,
			actual.Next)
		}
		if (actual.Previous != expected.Previous && (actual.Previous == nil  || expected.Previous == nil)) {
			
			t.Errorf(
`
-------- TEST FAILED --------
expected Previous: %v
actual Previous: %v
`,
			expected.Previous,
			actual.Previous)
		}
		if len(actual.Results) != len(expected.Results) {
			t.Errorf(
`
-------- TEST FAILED --------
expected len(actual.Results): %d
actual len(expected.Results): %d
`,
			len(actual.Results),
			len(expected.Results))
		}

		for i := range actual.Results {
			actualLocation := actual.Results[i]
			expectedLocation := expected.Results[i]

			if actualLocation.Name != expectedLocation.Name {
				t.Errorf(
`
-------- TEST FAILED --------
expected name: %s
actual name: %s
`,
				expectedLocation.Name,
				actualLocation.Name)
			}
			if actualLocation.URL != expectedLocation.URL {
				t.Errorf(
`
-------- TEST FAILED --------
expected URL: %s
actual URL: %s
`,
				expectedLocation.URL,
				actualLocation.URL)
			}
		}
	}

	client.StopCacheReap()
}