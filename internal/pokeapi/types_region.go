package pokeapi

type Region struct {
	ID        int `json:"id"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"locations"`
	MainGeneration struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"main_generation"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

type RegionList struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
