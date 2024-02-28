package pokeapi

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
	pokemon := Pokemon{}
	err := c.GetJson(baseURL+"/pokemon/"+idOrName, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
