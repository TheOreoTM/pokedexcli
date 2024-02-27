package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Not a valid command")
			continue
		}

		err := command.run(conf)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}

type Command struct {
	name        string
	description string
	run         func(*config) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			run:         runHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			run:         runExit,
		},
		"map": {
			name:        "map",
			description: "Go forward a page in the map",
			run:         runMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back a page on the map",
			run:         runMapB,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
