package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/osirisofficial/pokedex-cli/pokecache"
)

func main() {
	//=========  commands
	commands := map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": cliCommand{
			name:        "help",
			description: "help in using the pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "to get map area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "go back in map",
			callback:    commandMapd,
		},
		"explore": {
			name:        "explore",
			description: "explore pokemon at a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: " to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "to get details of caught pokemon",
			callback:    commandInspect,
		},
	}
	// to store the url for map & dmap command
	cf := config{
		next:     "",
		previous: "",
		count:    0,
	}

	// cache
	cache_map := pokecache.NewCache(15 * time.Second)

	// catch pokemon
	catchPokemon := make(map[string]CaughtPokemonDetails)

	//==========  REPL

	// create scanner
	scanner := bufio.NewScanner(os.Stdin)

	//infinite loop - loop(L) =============================
	for {
		// print starting
		fmt.Print("pokedex> ")

		// scan input - read(R) =============================
		if !scanner.Scan() { // read line
			break
		}

		// get the scan input as string
		input := scanner.Text()

		// cleaning the input
		clean_input := cleanInput(input)

		// check the input cmd & eval - eval(E) & print(P) ===============
		ele, ok := commands[clean_input[0]]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			if len(clean_input) > 1 {
				ele.callback(&cf, &cache_map, clean_input[1], catchPokemon)
			} else {
				ele.callback(&cf, &cache_map, "", catchPokemon)
			}
		}

	}

}
