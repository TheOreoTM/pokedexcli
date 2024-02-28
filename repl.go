package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
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
			name:        "map {f|b}",
			description: "Go forward a page in the map",
			aliases:     []string{"m"},
			run:         runMap,
		},
		"explore": {
			name:        "explore {area}",
			description: "View more information about an area",
			aliases:     []string{"ex", "e"},
			run:         runExplore,
		},
		"location": {
			name:        "location {id|name}",
			description: "View information about a location",
			aliases:     []string{"loc", "l"},
			run:         runLocation,
		},
		"region": {
			name:        "region {id|name}",
			description: "View information about a region",
			aliases:     []string{"reg", "r"},
			run:         runRegion,
		},
		"pokemons": {
			name:        "pokemons {f|b}",
			description: "View a list of pokemons",
			aliases:     []string{"pok", "pp"},
			run:         runPokemons,
		},
		"pokemon": {
			name:        "pokemon {id|name}",
			description: "View information about a pokemon",
			aliases:     []string{"p"},
			run:         runPokemon,
		},
		"catch": {
			name:        "catch {id|name}",
			description: "Catch a pokemon",
			aliases:     []string{"c"},
			run:         runCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your pokedex",
			aliases:     []string{"d"},
			run:         runPokedex,
		},
	}
}

func getCommand(nameOrAlias string) Command {
	commands := getCommands()

	for _, command := range commands {
		if strings.Split(command.name, " ")[0] == nameOrAlias {
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

func calculatePages(nextUrl *string, totalItems int) (int, int) {
	if nextUrl == nil {
		return 1, (totalItems-1)/20 + 1
	}

	u, err := url.Parse(*nextUrl)
	if err != nil {
		return 1, (totalItems-1)/20 + 1
	}

	q := u.Query()
	offset, err := strconv.Atoi(q.Get("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 20
	}

	currentPage := offset / limit
	if offset%limit != 0 {
		currentPage++
	}

	totalPages := totalItems / limit
	if totalItems%limit != 0 {
		totalPages++
	}

	return currentPage + 1, totalPages
}
