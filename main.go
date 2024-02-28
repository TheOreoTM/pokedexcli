package main

import (
	"time"

	"github.com/theoreotm/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	nextPokemonPageUrl  *string
	prevPokemonPageUrl  *string
	caughtPokemon       map[int]pokeapi.Pokemon
}

func main() {
	conf := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[int]pokeapi.Pokemon),
	}

	startRepl(&conf)
}
