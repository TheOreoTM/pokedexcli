package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullUrl := baseURL + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// check the cache
	if val, ok := c.cache.Get(fullUrl); ok {
		locationAreasRes := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationAreasRes)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasRes, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v: ", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasRes := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreasRes, nil
}
