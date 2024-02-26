package main

import (
	"time"

	"github.com/neetusuthar90/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocURL    *string
	prevLocURL    *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	StartRepl(&cfg)
}
