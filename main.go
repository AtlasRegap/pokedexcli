package main

import (
	"time"

	"github.com/AtlasRegap/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		cache:         pokecache.NewCache(5 * time.Minute),
		caughtPokemon: make(map[string]Pokemon),
	}
	startRepl(cfg)
}
