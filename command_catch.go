package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

func runCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon to catch")
	}

	pokemonID := strings.Join(args, "+")
	pokemon, err := conf.pokeapiClient.GetPokemon(pokemonID)
	if err != nil {
		return err
	}

	fmt.Printf("You throw a pokeball at the %s\n", pokemon.Name)

	const threshold = 50

	randNum := rand.IntN(pokemon.BaseExperience)

	if randNum < threshold {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("click... ")
		time.Sleep(750 * time.Millisecond)
		fmt.Printf("click... ")
		time.Sleep(1250 * time.Millisecond)
		fmt.Printf("click... \n")
		fmt.Printf("You caught the %s!\n", pokemon.Name)
	} else {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("click... ")
		time.Sleep(750 * time.Millisecond)
		fmt.Printf("click... ")
		time.Sleep(1250 * time.Millisecond)
		fmt.Printf("The %s broke free!\n", pokemon.Name)
	}

	return nil

}
