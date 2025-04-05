package main

import (
	"log"
	"os"
	"strings"

	pokecache "github.com/ziyad678/pokedexcli/internal"
)

func main() {
	cfg := config{}
	cfg.Commands = cfg.initCommands()
	cfg.LocCache = pokecache.NewCache(15)
	cfg.Pokemons = map[string]Pokemon{}
	logFileName := "app.log"
	//logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	cfg.repl()

}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))
}
