package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are available commands:")

	availableCommand := getCommands()

	for _, cmd := range availableCommand {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
