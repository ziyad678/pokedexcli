package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
