package main

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
		"explore": {
			name:        "explore",
			description: "List pokemons in a given area",
			callback:    commandExplore,
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
		"inspect": {
			name:        "inspect",
			description: "List pokemon stats",
			callback:    commandInspect,
		},
	}

	return commands
}
