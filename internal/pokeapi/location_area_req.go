package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullUrl := baseURL + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	locationAreasRes := LocationAreasResponse{}
	err := c.GetJson(fullUrl, &locationAreasRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasRes, nil
}

func (c *Client) GetLocationArea(idOrName string) (LocationArea, error) {
	endpoint := fmt.Sprintf("/location-area/%s", idOrName)
	fullUrl := baseURL + endpoint

	locationArea := LocationArea{}
	err := c.GetJson(fullUrl, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	bytes, err := json.Marshal(locationArea)

	if err == nil {
		c.cache.Add(fmt.Sprintf("%s/location-area/%s", baseURL, locationArea.Name), bytes)
		c.cache.Add(fmt.Sprintf("%s/location-area/%d", baseURL, locationArea.ID), bytes)
	}

	return locationArea, nil
}
