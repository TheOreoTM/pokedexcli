package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetRegion(idOrName string) (Region, error) {
	var region Region
	err := c.GetJson(baseURL+"/region/"+idOrName, &region)
	if err != nil {
		return Region{}, err
	}

	bytes, err := json.Marshal(region)

	if err == nil {
		c.cache.Add(fmt.Sprintf("%s/region/%s", baseURL, region.Name), bytes)
		c.cache.Add(fmt.Sprintf("%s/region/%d", baseURL, region.ID), bytes)
	}

	return region, nil
}

func (c *Client) GetRegionList() (RegionList, error) {
	var regionList RegionList
	err := c.GetJson(baseURL+"/region", &regionList)
	if err != nil {
		return RegionList{}, err
	}

	return regionList, nil
}
