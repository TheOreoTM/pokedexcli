package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetJson sends a GET request to the specified URL and parses the JSON response into the provided interface.
// It returns an error if the request fails or if the response has a bad status code.
// Additionally, it adds the response data to the cache.
func (c *Client) GetJson(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		if res.StatusCode == 404 {
			return fmt.Errorf("not found")
		}
		return fmt.Errorf("bad status code: %v: ", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	c.cache.Add(url, data)

	return nil
}
