package main

import pokecache "github.com/ziyad678/pokedexcli/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
type config struct {
	Next     *string
	Previous *string
	Commands map[string]cliCommand
	LocCache pokecache.Cache
	Pokemons map[string]Pokemon
}

type Location struct {
	Name string `json:"name"`
}
