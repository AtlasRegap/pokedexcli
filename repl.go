package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AtlasRegap/pokedexcli/internal/pokecache"
)

type config struct {
	commands         map[string]cliCommand
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
	caughtPokemon    map[string]Pokemon
}





func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
