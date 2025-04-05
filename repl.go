package main

import (
	"bufio"
	"fmt"
	"os"
)

func (cfg *config) repl() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		cleaned := cleanInput(sc.Text())
		command, ok := cfg.Commands[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Print("Pokedex > ")
			continue
		}
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		err := command.callback(cfg,args...)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Pokedex > ")
	}
}
