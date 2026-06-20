package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (RespPokemons, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespPokemons{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespPokemons{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemons{}, err
	}

	c.cache.Add(url, dat)

	locationResp := RespPokemons{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespPokemons{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil

}
