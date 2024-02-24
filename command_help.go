package main

import "fmt"

func runHelp() {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available commands")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("- %v\n", command.name)
	}

	fmt.Println("")
}
