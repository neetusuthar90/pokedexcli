package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := CleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommand := getCommands()

		command, ok := availableCommand[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List of next location areas:",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List of previous location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "pokemon in your pokedex",
			callback:    callbackPokedex,
		},
		"inspect": {
			name:        "inspect {inspect_name}",
			description: "Inspect the caught pokemon",
			callback:    callbackInspect,
		},
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
