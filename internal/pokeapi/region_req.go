package pokeapi

func (c *Client) GetRegion(idOrName string) (Region, error) {
	var region Region
	err := c.GetJson(baseURL+"/region/"+idOrName, &region)
	if err != nil {
		return Region{}, err
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
