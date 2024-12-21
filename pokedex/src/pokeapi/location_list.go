package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageUrl *string) (Locations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := client.cache.Get(url); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Locations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, nil
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Locations{}, fmt.Errorf("can't get locations: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	var locations Locations
	if err = json.Unmarshal(data, &locations); err != nil {
		return Locations{}, err
	}

	client.cache.Add(url, data)
	return locations, nil
}
