package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		return errors.New("You don't have any pokemons. Catch it first")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
