package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocation(idOrName string) (Location, error) {
	url := fmt.Sprintf("%s/location/%s", baseURL, idOrName)

	var location Location
	err := c.GetJson(url, &location)
	if err != nil {
		return Location{}, err
	}

	bytes, err := json.Marshal(location)

	if err == nil {
		c.cache.Add(fmt.Sprintf("%s/location/%s", baseURL, location.Name), bytes)
		c.cache.Add(fmt.Sprintf("%s/location/%d", baseURL, location.ID), bytes)
	}

	return location, nil
}
