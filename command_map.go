package main

import (
	"errors"
	"fmt"
	"strings"
)

func runMap(conf *config, args ...string) error {
	if len(args) == 0 {
		return runMapNext(conf)
	}

	switch args[0] {
	case "f":
		return runMapNext(conf, args[1:]...)
	case "b":
		return runMapPrev(conf, args[1:]...)
	default:
		return runMapNext(conf, args...)
	}
}

func runMapNext(conf *config, _ ...string) error {
	areas, err := conf.pokeapiClient.ListLocationAreas(conf.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Printf("Location areas:\n")
	for _, area := range areas.Results {
		splits := strings.Split(area.URL, "/")
		id := splits[len(splits)-2] // Because there is a slash at the end

		fmt.Printf("%s. %s\n", id, area.Name)
	}

	conf.nextLocationAreaURL = areas.Next
	conf.prevLocationAreaURL = areas.Previous

	return nil
}

func runMapPrev(conf *config, _ ...string) error {
	if conf.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	areas, err := conf.pokeapiClient.ListLocationAreas(conf.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Printf("Location areas:\n")
	for _, area := range areas.Results {
		splits := strings.Split(area.URL, "/")
		id := splits[len(splits)-2] // Because there is a slash at the end

		fmt.Printf("%s. %s\n", id, area.Name)
	}

	conf.nextLocationAreaURL = areas.Next
	conf.prevLocationAreaURL = areas.Previous

	return nil
}
