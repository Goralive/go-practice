package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New("you have not caught the pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%s \n", typeInfo.Type.Name)
	}
	return nil
}
