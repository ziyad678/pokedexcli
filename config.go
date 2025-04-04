package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
type config struct {
	Next     string
	Previous string
	Commands map[string]cliCommand
}

type Location struct {
	Name string `json:"name"`
}
