package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/AtlasRegap/pokedexcli/internal/pokecache"
	"github.com/peterh/liner"
)

type config struct {
	commands         map[string]cliCommand
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
	caughtPokemon    map[string]Pokemon
	line             *liner.State
}

func startRepl(cfg *config) {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	cfg.line = line

	for {
		input, err := line.Prompt("Pokedex > ")
		if err == liner.ErrPromptAborted || err == io.EOF {
			fmt.Println()
			return
		}
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		line.AppendHistory(input)

		commandName := words[0]
		args := words[1:]

		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err = command.callback(cfg, args...)
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
