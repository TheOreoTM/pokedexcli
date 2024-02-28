package pokeapi

import (
	"fmt"
)

func (c *Client) GetLocation(idOrName string) (Location, error) {
	url := fmt.Sprintf("%s/location/%s", baseURL, idOrName)

	var location Location
	err := c.GetJson(url, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}
