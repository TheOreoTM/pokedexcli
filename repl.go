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

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Not a valid command")
			continue
		}

		err := command.run(conf, args...)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}

type Command struct {
	name        string
	description string
	run         func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "View more information about an area",
			run:         runExplore,
		},
		"location": {
			name:        "location",
			description: "View information about a location",
			run:         runLocation,
		},
		"region": {
			name:        "region",
			description: "View information about a region",
			run:         runRegion,
		},
		"pokemons": {
			name:        "pokemons",
			description: "View a list of pokemons",
			run:         runPokemons,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
