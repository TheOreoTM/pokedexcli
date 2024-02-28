package main

import (
	"fmt"
	"strings"
)

func runPokemons(conf *config, args ...string) error {
	if len(args) == 0 {
		return runPokemonsForward(conf)
	}

	switch args[0] {
	case "f":
		return runPokemonsForward(conf, args[1:]...)
	case "b":
		return runPokemonsBackward(conf, args[1:]...)
	default:
		return runPokemonsForward(conf, args...)
	}
}

func runPokemonsForward(conf *config, _ ...string) error {
	pokemons, err := conf.pokeapiClient.GetPokemons(conf.nextPokemonPageUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons:\n")
	maxIDLength := 0
	for _, pokemon := range pokemons.Results {
		id := strings.Split(pokemon.URL, "/")
		idStr := id[len(id)-2]
		if len(idStr) > maxIDLength {
			maxIDLength = len(idStr)
		}
	}

	for _, pokemon := range pokemons.Results {
		id := strings.Split(pokemon.URL, "/")
		idStr := id[len(id)-2]
		fmt.Printf("%*s. %s\n", maxIDLength, idStr, pokemon.Name)
	}

	conf.prevPokemonPageUrl = pokemons.Previous
	conf.nextPokemonPageUrl = pokemons.Next

	return nil
}

func runPokemonsBackward(conf *config, _ ...string) error {
	if conf.prevPokemonPageUrl == nil {
		return fmt.Errorf("you're on the first page")
	}

	pokemons, err := conf.pokeapiClient.GetPokemons(conf.prevPokemonPageUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons:\n")
	maxIDLength := 0
	for _, pokemon := range pokemons.Results {
		id := strings.Split(pokemon.URL, "/")
		idStr := id[len(id)-2]
		if len(idStr) > maxIDLength {
			maxIDLength = len(idStr)
		}
	}

	for _, pokemon := range pokemons.Results {
		id := strings.Split(pokemon.URL, "/")
		idStr := id[len(id)-2]
		fmt.Printf("%*s. %s\n", maxIDLength, idStr, pokemon.Name)
	}

	conf.prevPokemonPageUrl = pokemons.Previous
	conf.nextPokemonPageUrl = pokemons.Next

	return nil
}
