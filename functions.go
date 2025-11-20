package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/osirisofficial/pokedex-cli/pokecache"
)

// function to clean input
// example : "   manas  konda   "
// result : []string{"manas","konda"}
func cleanInput(text string) []string {
	text = strings.ToLower(text) // to lower case
	var word string
	var res []string
	var len = len(text) - 1
	// to gather words in slice and remove unwanted  spaces
	for indx, ele := range strings.Split(text, "") {

		if ele != " " { // hello
			word = fmt.Sprintf("%s%s", word, ele)
		}

		if ele == " " && word != "" { // __hello_
			res = append(res, word)
			word = ""
		}

		if indx == len && word != "" {
			res = append(res, word)

		}

	}
	return res
}

// function to end pokedex on use "exit"
func commandExit(cf *config, cache_map *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {
	fmt.Println("Closing the Pokendex... Goodbye!")
	os.Exit(0)
	return nil
}

// function on use "help"
func commandHelp(cf *config, cache_map *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

// function on use "map"
func commandMap(cf *config, cache *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {

	//update urls in config
	(*cf).count++
	if (*cf).count == 1 {
		(*cf).next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", (*cf).count-1)
		(*cf).previous = (*cf).next
	}
	if (*cf).count != 1 {
		(*cf).next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count-1)*20)
		(*cf).previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count-2)*20)
	}
	fmt.Println("map : ", cf)

	// making request or using cache if present
	body, status := request_and_cache((*cf).next, cache)

	if status != nil {
		errors.New("check you connection")
	}
	// converting data from json []byte to struct and priniting it
	fmt.Println(unmarshal_result(body))

	return nil
}

// function on use "dmap"
func commandMapd(cf *config, cache *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {

	body, status := request_and_cache((*cf).previous, cache)

	if status != nil {
		errors.New("check you connection")
	}

	fmt.Println(unmarshal_result(body))

	////update urls in config

	if (*cf).count > 1 {
		(*cf).count--
		(*cf).previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count-1)*20)
		(*cf).next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count+1)*20)
	} else {
		(*cf).next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count+1)*20)
		(*cf).previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", ((*cf).count-1)*20)

	}

	//fmt.Println("mapb : ", cf)
	return nil
}

// func on use "explore {argument}"
func commandExplore(cf *config, cache *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", input)

	body, err := request_and_cache(url, cache)

	if err != nil {
		return err
	}

	var result PokemonEncounterResponse
	json.Unmarshal(body, &result)

	fmt.Println(result)
	return nil
}

// func on use "catch"
func commandCatch(cf *config, cache *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {
	fmt.Println("Throwing a Pokeball at ", input)

	response, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", input))

	//handel error
	if err != nil {
		return err
	}

	//closing response body
	defer response.Body.Close()

	//reading the response body
	body, err := io.ReadAll(response.Body)

	//unmarshal body
	var res CaughtPokemonDetails
	json.Unmarshal(body, &res)

	//adding pokemon to caughtPokemon
	caughtPokemon[input] = res

	fmt.Println(input, " was caught!")

	return nil

}

// func on use inspect
func commandInspect(cf *config, cache *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error {
	val, ok := caughtPokemon[input]

	if !ok {
		fmt.Println("pokemon not caught")
		return nil
	}

	fmt.Println(val)

	return nil
}

// to make new request or use cache if present
func request_and_cache(url string, cache_map *pokecache.Cache) ([]byte, error) {

	/////////using cache data
	data, status := cache_map.Get(url)
	if status {
		return data, nil
	}

	///////// making request data does not exists in cache
	//get request
	response, err := http.Get(url)

	//handel error
	if err != nil {
		return nil, err
	}

	//closing response body
	defer response.Body.Close()

	//reading the response body
	body, err := io.ReadAll(response.Body)

	// adding data to cache
	cache_map.Add(url, body)

	return body, nil
}

// convert []byte into struct for "map" and "mapd"
func unmarshal_result(body []byte) LocationAreaList {
	var body_result LocationAreaList
	json.Unmarshal(body, &body_result)

	return body_result
}
