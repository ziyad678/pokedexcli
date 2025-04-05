package main

import (
	"errors"
	"fmt"
	"log"
)

func commandInspect(cfg *config, args ...string) error {
	log.Println("Entering Inspect Function")
	if len(args) < 1 {
		log.Println("No pokemon name provided to inspect. returning")
		return errors.New("please enter a pokemon name to inspect")
	}

	pokemon, ok := cfg.Pokemons[args[0]]
	if !ok {
		log.Printf("%v was not caught yet. can't be inspected\n", args[0])
		return errors.New("pokemon not caught yet")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Stats:")

	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
