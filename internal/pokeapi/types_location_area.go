package pokeapi

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Result   []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"result"`
}
