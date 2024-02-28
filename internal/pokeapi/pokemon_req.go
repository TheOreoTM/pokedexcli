package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemons(pageUrl *string) (PokemonsResponse, error) {
	endpoint := "/pokemon?offset=0&limit=20"
	fullUrl := baseURL + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	pokemons := PokemonsResponse{}
	err := c.GetJson(fullUrl, &pokemons)
	if err != nil {
		return PokemonsResponse{}, err
	}

	return pokemons, nil
}

func (c *Client) GetPokemon(idOrName string) (Pokemon, error) {
	if idOrName == "" {
		return Pokemon{}, fmt.Errorf("you must provide a pokemon")
	}

	pokemon := Pokemon{}
	err := c.GetJson(baseURL+"/pokemon/"+idOrName, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	bytes, err := json.Marshal(pokemon)

	if err == nil {
		c.cache.Add(fmt.Sprintf("%s/pokemon/%s", baseURL, pokemon.Name), bytes)
		c.cache.Add(fmt.Sprintf("%s/pokemon/%d", baseURL, pokemon.ID), bytes)
	}

	return pokemon, nil
}
