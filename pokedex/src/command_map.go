package main

import (
	"fmt"

	"errors"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = locationsResp.Next
	cfg.prevLocationUrl = locationsResp.Previous
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = locationsResp.Next
	cfg.prevLocationUrl = locationsResp.Previous
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
