package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Method with receiver are used for interface purpose
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache find
		fmt.Println("Cache hit!!")
		locAreaResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locAreaResp, nil
	}
	fmt.Println("Cache missing!!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResp{}, err
	}

	locAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullURL, data)
	return locAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache find
		fmt.Println("Cache hit!!")
		locArea := LocationArea{}
		err := json.Unmarshal(data, &locArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locArea, nil
	}
	fmt.Println("Cache missing!!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locArea := LocationArea{}
	err = json.Unmarshal(data, &locArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, data)
	return locArea, nil
}
