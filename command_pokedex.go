package main

import (
	"fmt"
)

func runPokedex(conf *config, _ ...string) error {
	caughtPokemons := conf.caughtPokemon
	if len(caughtPokemons) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		fmt.Println("Use the `catch` command to catch a pokemon.")
		return nil
	}

	fmt.Println("Your caught pokemon:")
	for _, pokemon := range caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
