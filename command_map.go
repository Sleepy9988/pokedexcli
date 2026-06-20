package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args ...string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)

	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}

func commandMapBack(c *config, args ...string) error {

	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)

	if err != nil {
		return err
	}

	c.nextLocationsURL = locationResp.Next
	c.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
