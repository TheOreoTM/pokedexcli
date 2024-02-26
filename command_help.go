package main

import "fmt"

func runHelp() {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available commands")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	fmt.Println("")
}
