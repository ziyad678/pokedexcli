package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
type config struct {
	Next     string
	Previous string
}

type Location struct {
	Name string `json:"name"`
}

func main() {
	cfg := config{}
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
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		cleaned := cleanInput(sc.Text())
		command, ok := commands[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Print("Pokedex > ")
			continue
		}
		command.callback()
		fmt.Print("Pokedex > ")
	}

}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))
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
