package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args ...string) error {
	log.Println("Entering Catch function")
	if len(args) < 1 {
		log.Println("No pokemon name provided to catch. returning")
		return errors.New("please enter a pokemon name to catch")
	}
	pokemonResult, err := cfg.getPokemon(args[0])
	if err != nil {
		log.Printf("GetPokemon return an errorr for %v. Error: %v\n", args[0], err)
		return err
	}
	res := rand.Intn(pokemonResult.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResult.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonResult.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResult.Name)
	cfg.Pokemons[pokemonResult.Name] = pokemonResult

	return nil
}

func (cfg *config) getPokemon(name string) (PokemonResponse, error) {
	// Define the API endpoint URL
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	log.Println("Inside getPokemon, going to hit", url)
	log.Printf("Checking if entry for %v exists in cache\n", url)
	var apiResponse PokemonResponse
	cacheEntry, found := cfg.LocCache.Get(url)
	if found {
		log.Printf("Found entry in cache for %v\n", url)
		err := json.Unmarshal(cacheEntry, &apiResponse)
		if err != nil {
			log.Printf("Error unmarshaling JSON from cache entry: %v", err)
			return PokemonResponse{}, err
		}
		log.Printf("Returning cache entry for %v\n", url)
		return apiResponse, nil
	}
	log.Printf("No cache entry found for %v. Initiating Get request\n", url)
	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", url, err)
		return PokemonResponse{}, err
	}

	defer resp.Body.Close()

	// Check if the request was successful (status code 200 OK)
	if resp.StatusCode == http.StatusNotFound {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Error: Received non-404 status code: %d\nResponse Body: %s", resp.StatusCode, string(bodyBytes))
		return PokemonResponse{}, errors.New("invalid Pokemon name. Please enter correct Pokemon name")
	}
	if resp.StatusCode != http.StatusOK {
		// Read the body even on error for potential error messages from the API
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Error: Received non-200 status code: %d\nResponse Body: %s", resp.StatusCode, string(bodyBytes))
		return PokemonResponse{}, errors.New("failed to get 200 respone to API call")
	}
	log.Printf("Reading response body for %v\n", url)
	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return PokemonResponse{}, err
	}
	//add to cache
	log.Printf("Adding entry for %v in cache\n", url)
	cfg.LocCache.Add(url, bodyBytes)

	// Unmarshal (parse) the JSON byte slice into the Go struct
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return PokemonResponse{}, err
	}
	log.Printf("Leaving explroeLocation for %v\n", url)
	return apiResponse, nil
}
