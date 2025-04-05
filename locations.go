package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokemonLocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Struct to represent the overall JSON response structure from the API
type PokemonLocationAreaResponse struct {
	Count    int                   `json:"count"`
	Next     *string               `json:"next"`     // Use pointer for nullable fields
	Previous *string               `json:"previous"` // Use pointer for nullable fields
	Results  []PokemonLocationArea `json:"results"`
}

func (cfg *config) listLocations(pageURL *string) (PokemonLocationAreaResponse, error) {
	// Define the API endpoint URL
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	var apiResponse PokemonLocationAreaResponse
	if cfg.Next != nil {
		url = *pageURL
	}
	log.Println("Inside listLocation, going to hit",url)
	log.Printf("Checking if entry for %v exists in cache\n",url)
	cacheEntry, found := cfg.LocCache.Get(url)

	if found {
		log.Printf("Found entry in cache for %v\n",url)
		err := json.Unmarshal(cacheEntry, &apiResponse)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON from cache entry: %v", err)
			return PokemonLocationAreaResponse{}, err
		}
		log.Printf("Returning cache entry for %v\n",url)
		return apiResponse, nil
	}
	log.Printf("No cache entry found for %v. Initiating Get request\n",url)
	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL %s: %v", url, err)
		return PokemonLocationAreaResponse{}, err
	}
	// Ensure the response body is closed even if errors occur later
	defer resp.Body.Close()

	// Check if the request was successful (status code 200 OK)
	if resp.StatusCode != http.StatusOK {
		// Read the body even on error for potential error messages from the API
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Fatalf("Error: Received non-200 status code: %d\nResponse Body: %s", resp.StatusCode, string(bodyBytes))
		return PokemonLocationAreaResponse{}, err
	}
	log.Printf("Reading response body for %v\n",url)
	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return PokemonLocationAreaResponse{}, err
	}
	//add to cache
	log.Printf("Adding entry for %v in cache\n",url)
	cfg.LocCache.Add(url, bodyBytes)

	// Unmarshal (parse) the JSON byte slice into the Go struct
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return PokemonLocationAreaResponse{}, err
	}
	log.Printf("Leaving listLocation for %v\n",url)
	return apiResponse, nil
}
