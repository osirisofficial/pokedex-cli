package main

import "github.com/osirisofficial/pokedex-cli/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(cf *config, cache_map *pokecache.Cache, input string, caughtPokemon map[string]CaughtPokemonDetails) error
	calls       int
}

// to store urls
type config struct {
	next     string
	previous string
	count    int
}

// to convert json to struct
type LocationAreaList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaItem `json:"results"`
}

type LocationAreaItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounterResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CaughtPokemonDetails struct {
	Height         int                         `json:"height"`
	Weight         int                         `json:"weight"`
	BaseExperience int                         `json:"base_experience"`
	Stats          []CaughtPokemonDetailsStats `json:"stats"`
	Types          []CaughtPokemonDetailsTypes `json:"types"`
}

type CaughtPokemonDetailsStats struct {
	BaseStat int                           `json:"base_stat"`
	StatName CaughtPokemonDetailsStatsName `json:"stat"`
}

type CaughtPokemonDetailsStatsName struct {
	Name string `json:"name"`
}

type CaughtPokemonDetailsTypes struct {
	Type CaughtPokemonDetailsTypesName
}

type CaughtPokemonDetailsTypesName struct {
	Name string `json:"name"`
}
