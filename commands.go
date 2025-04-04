package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (cfg *config) initCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cfg.commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help menu",
			callback:    cfg.commandHelp,
		},
		"map": {
			name:        "map",
			description: "List 20 location areas",
			callback:    cfg.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas",
			callback:    cfg.commandMapBack,
		},
	}
	return commands
}

func (cfg *config) commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func (cfg *config) commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func (cfg *config) commandMap() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Println("Can't parse URL")
	}
	defer resp.Body.Close()
	locations := Location{}
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(locations)
	return nil
}

func (cfg *config) commandMapBack() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
