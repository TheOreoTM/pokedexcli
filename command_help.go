package main

import (
	"fmt"
	"strings"
)

func runHelp(conf *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available commands")

	commands := getCommands()

	for _, command := range commands {
		prefix := fmt.Sprintf(" - %s: ", command.name)
		fmt.Printf("%s%s\n%*sAliases: %s\n\n", prefix, command.description, len(prefix), "", strings.Join(command.aliases, ", "))
	}

	fmt.Println("")

	return nil
}
