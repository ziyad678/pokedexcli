package main

import (
	"strings"
)

func main() {
	cfg := config{}
	cfg.Commands = cfg.initCommands()
	cfg.repl()

}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))
}
