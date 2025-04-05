package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func (cfg *config) initCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help menu",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List 20 location areas",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas",
			callback:    commandMapBack,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List caught pokemons",
			callback:    commandPokedex,
		},
	}

	return commands
}

func commandExit(cfg  *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func commandMapForward(cfg *config, args...string) error {
	log.Println("Map Forward called, calling ListLocations for ", cfg.Next)

	locationResponse, err := cfg.listLocations(cfg.Next)
	if err != nil {
		log.Println("listLocation returned an error", err)
		return err
	}
	log.Printf("Updating Next and Previous after succsessful listLocation call. Before updating values are Next:%v, Previous:%v\n", cfg.Next,cfg.Previous)
	cfg.Next = locationResponse.Next
	cfg.Previous = locationResponse.Previous
	log.Printf("Next and Previous updated. New values are Next:%v, Previous:%v\n", cfg.Next,cfg.Previous)
	log.Println("printing list of locations")
	for _, area := range locationResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args...string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	locationResponse, err := cfg.listLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	cfg.Previous = locationResponse.Previous

	for _, area := range locationResponse.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	cfg.Pokemons = append(cfg.Pokemons, args[0])
	return nil
}

func commandPokedex(cfg *config, args...string) error {
	for _,v := range cfg.Pokemons{
		fmt.Println(v)
	}

	return nil
}

