package main

import (
	"errors"
	"fmt"
	"strings"
)

func runPokemons(conf *config, args ...string) error {
	if len(args) == 0 {
		return runPokemonsNext(conf)
	}

	switch args[0] {
	case "f":
		return runPokemonsNext(conf, args[1:]...)
	case "b":
		return runPokemonsPrev(conf, args[1:]...)
	default:
		return runPokemonsNext(conf, args[1:]...)
	}
}

func runPokemon(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon id")
	}

	pokemonID := strings.Join(args, "+")
	pokemon, err := conf.pokeapiClient.GetPokemon(pokemonID)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Abilities:\n")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil

}

func runPokemonsNext(conf *config, _ ...string) error {
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

	currentPage, totalPages := calculatePages(conf.nextPokemonPageUrl, pokemons.Count)
	conf.prevPokemonPageUrl = pokemons.Previous
	conf.nextPokemonPageUrl = pokemons.Next

	fmt.Printf("Page %d of %d\n", currentPage, totalPages)

	return nil
}

func runPokemonsPrev(conf *config, _ ...string) error {
	if conf.prevPokemonPageUrl == nil {
		return errors.New("you're on the first page")
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

	currentPage, totalPages := calculatePages(conf.prevPokemonPageUrl, pokemons.Count)
	conf.prevPokemonPageUrl = pokemons.Previous
	conf.nextPokemonPageUrl = pokemons.Next

	fmt.Printf("Page %d of %d\n", currentPage, totalPages)

	return nil
}
