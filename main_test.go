package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/osirisofficial/pokedex-cli/pokecache"
)

func TestCleanInput(t *testing.T) {
	fmt.Println("done")
	//test cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ", // "_", "_hello" "_world" "_"
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world manas", // "_", "_hello" "_world" "_"
			expected: []string{"hello", "world", "manas"},
		},
	}

	// then loop over to run test case
	for _, c := range cases {
		// calling function with test case as paramter
		actual := cleanInput(c.input)

		//matching len of output with exepcted output
		if len(actual) != len(c.expected) {
			t.Errorf("test failed")
		}

		// if len match check output with expected output
		for i, ele := range actual {
			expected_word := c.expected[i]

			if ele != expected_word {
				t.Errorf("test failed")
			}
		}

	}
}

func TestExplore(t *testing.T) {
	cases := []string{"canalave-city-area", "eterna-city-area"}
	cache_map := pokecache.NewCache(15 * time.Second)

	for _, ele := range cases {
		commandExplore(nil, &cache_map, ele)
	}
}

func TestCatch(t *testing.T) {
	cases := []string{"tentacool", "tentacruel", "staryu", "magikarp"}
	cache_map := pokecache.NewCache(15 * time.Second)
	caughtPokemon := map[string]CaughtPokemonDetails{}

	for _, ele := range cases {
		commandCatch(nil, &cache_map, ele, caughtPokemon)
	}
}
