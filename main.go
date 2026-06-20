package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}
