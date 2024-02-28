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
