package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMapForward(cfg *config, args ...string) error {
	log.Println("Map Forward called, calling ListLocations for ", cfg.Next)

	locationResponse, err := cfg.listLocations(cfg.Next)
	if err != nil {
		log.Println("listLocation returned an error", err)
		return err
	}
	log.Printf("Updating Next and Previous after succsessful listLocation call. Before updating values are Next:%v, Previous:%v\n", cfg.Next, cfg.Previous)
	cfg.Next = locationResponse.Next
	cfg.Previous = locationResponse.Previous
	log.Printf("Next and Previous updated. New values are Next:%v, Previous:%v\n", cfg.Next, cfg.Previous)
	log.Println("printing list of locations")
	for _, area := range locationResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
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
