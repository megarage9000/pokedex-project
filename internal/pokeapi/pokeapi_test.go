package pokeapi

import (
	"testing"
	"time"
)

func createDefaultClient() Client {
	return 	NewClient(5 * time.Second, 1 * time.Minute)
}

func TestPokeApiMap(t *testing.T) {

	inputStr := "https://pokeapi.co/api/v2/location-area"
	nextInputStr := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

	cases := []struct {
		input string
		expected LocationArea
	} {
		{
			input: inputStr,
			expected: LocationArea {
				Count: 1089,
				Next: &nextInputStr,
				Previous: nil,
				Results: []LocationAreaResult {
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

	client := createDefaultClient()

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

func TestPokeApiExplore(t *testing.T) {

	cases := []struct {
		input string
		expected struct {
			name string
			pokemon []string
		}
	} {
		{
			input: "canalave-city-area",
			expected: struct {
				name string
				pokemon []string
			} {
				name: "canalave-city-area",
				pokemon: []string {
					"tentacool",
					"tentacruel",
					"staryu",
					"magikarp",
					"gyarados",
					"wingull",
					"pelipper",
					"shellos",
					"gastrodon",
					"finneon",
					"lumineon",
				},
			},
		},
	}

	client := createDefaultClient()
	
	for _, testCase := range cases {

		actual, err := client.ExploreLocation(testCase.input)

		if err != nil {
			t.Errorf(
`
-------- TEST FAILED --------
error when calling ExploreLocation(): %s
`, err)
		}
		
		expectedResults := testCase.expected
		actualName := actual.Name

		if expectedResults.name != actualName {
			t.Errorf(
`
-------- TEST FAILED --------
expected name on LocationPokemonList: %s
actual name on LocationPokemonList: %s
`, expectedResults.name, actualName)
		}

		if len(actual.PokemonEncounters) != len(expectedResults.pokemon) {
			t.Errorf(
`
-------- TEST FAILED --------
expected # of pokemon encounters on LocationPokemonList: %d
actual # of pokemon encounters on LocationPokemonList: %d
`, len(actual.PokemonEncounters), len(expectedResults.pokemon))
		}

		for i, pokemonEncounter := range actual.PokemonEncounters {

			actualPokemonName := pokemonEncounter.Pokemon.Name
			expectedPokemonName := expectedResults.pokemon[i]

			if actualPokemonName != expectedPokemonName {
				t.Errorf(
`
-------- TEST FAILED --------
expected pokemon encounter on LocationPokemonList: %s
actual pokemon encounter on LocationPokemonList: %s
`, expectedPokemonName, actualPokemonName)
			}
		}
	}
}

func TestPokeApiPokemon(t *testing.T) {
	cases := []struct {
		input string
		expected struct {
			name string
			weight int
			height int
			stats map[string]int
		}
	} {
		{
			input: "charmander",
			expected: struct {
				name string
				weight int
				height int
				stats map[string]int
			} {
				name: "charmander",
				weight: 85,
				height: 6,
				stats: map[string]int {
					"hp": 39,
					"attack": 52,
					"defense": 43,
					"special-attack": 60,
					"special-defense": 50,
					"speed": 65,
				},
			},
		},
	}

	client := createDefaultClient()

	for _, testCase := range cases {

		actual, err := client.GetPokemonInformation(testCase.input)
		expectedResults := testCase.expected

		if err != nil {
			t.Errorf(
`
-------- TEST FAILED --------
error when calling ExploreLocation(): %s
`, 			err)
		}

		if actual.Name != expectedResults.name {
			t.Errorf(
`
-------- TEST FAILED --------
expected name on Pokemon: %s
actual name on Pokemon: %s
`, 			expectedResults.name, actual.Name)			
		}

		if actual.Weight != expectedResults.weight {
			t.Errorf(
`
-------- TEST FAILED --------
expected weight on Pokemon: %d
actual weight on Pokemon: %d
`, 			expectedResults.weight, actual.Weight)					
		}

		if actual.Height != expectedResults.height {
			t.Errorf(
`
-------- TEST FAILED --------
expected height on Pokemon: %d
actual height on Pokemon: %d
`, 			expectedResults.height, actual.Height)					
		}

		for _, stat := range actual.Stats {
			statName := stat.Stat.Name
			statValue := stat.BaseStat

			if expectedResults.stats[statName] != statValue {
				t.Errorf(
`
-------- TEST FAILED --------
expected stat %s value on Pokemon: %d
actual stat %s value on Pokemon: %d
`, 				statName, expectedResults.stats[statName], statName, statValue)	
			}
		}
	}
}