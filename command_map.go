package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous
	return nil
}
