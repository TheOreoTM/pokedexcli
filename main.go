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
}

func main() {
	conf := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&conf)
}
