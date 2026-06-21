package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := c.caughtPokemon[name]
	if !ok {
		fmt.Print("you haven't caught this pokemon yet\n")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")

	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	return nil
}
