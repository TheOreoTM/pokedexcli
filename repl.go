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

		command := getCommand(commandName)

		if command.name == "" {
			fmt.Printf("Command not found: %s\n", commandName)
			continue
		}

		err := command.run(conf, args...)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

	}
}

type Command struct {
	name        string
	description string
	aliases     []string
	run         func(*config, ...string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			aliases:     []string{"h"},
			run:         runHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			aliases:     []string{"quit", "q"},
			run:         runExit,
		},
		"map": {
			name:        "map",
			description: "Go forward a page in the map",
			aliases:     []string{"m"},
			run:         runMap,
		},
		"explore": {
			name:        "explore",
			description: "View more information about an area",
			aliases:     []string{"ex", "e"},
			run:         runExplore,
		},
		"location": {
			name:        "location",
			description: "View information about a location",
			aliases:     []string{"loc", "l"},
			run:         runLocation,
		},
		"region": {
			name:        "region",
			description: "View information about a region",
			aliases:     []string{"reg", "r"},
			run:         runRegion,
		},
		"pokemons": {
			name:        "pokemons",
			description: "View a list of pokemons",
			aliases:     []string{"p", "pokemon"},
			run:         runPokemons,
		},
	}
}

func getCommand(nameOrAlias string) Command {
	commands := getCommands()

	for _, command := range commands {
		if command.name == nameOrAlias {
			return command
		}
		for _, alias := range command.aliases {
			if alias == nameOrAlias {
				return command
			}
		}
	}

	return Command{}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
