package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func runMap(conf *config) error {
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

func runMapB(conf *config) error {
	if conf.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	areas, err := conf.pokeapiClient.ListLocationAreas(conf.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
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
