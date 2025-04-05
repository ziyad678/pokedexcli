package main

import "fmt"

func commandExplore(cfg *config, args...string) error {
	for _,v := range cfg.Pokemons{
		fmt.Println(v)
	}

	return nil
}