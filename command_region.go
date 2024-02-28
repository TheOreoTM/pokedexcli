package main

import (
	"fmt"
	"strings"
)

func runRegion(conf *config, args ...string) error {
	if len(args) == 0 {
		return runRegionList(conf)
	}

	switch args[0] {
	case "list":
		return runRegionList(conf)
	case "show":
		return runRegionShow(conf, args[1:]...)
	default:
		return runRegionShow(conf, args...)
	}
}

func runRegionList(conf *config, _ ...string) error {
	regionList, err := conf.pokeapiClient.GetRegionList()
	if err != nil {
		return err
	}

	fmt.Println("Regions:")
	for _, region := range regionList.Results {
		fmt.Printf(" - %s\n", region.Name)
	}

	return nil
}

func runRegionShow(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Region ID or name")
		return nil
	}

	idOrName := strings.Join(args, " ")

	region, err := conf.pokeapiClient.GetRegion(idOrName)

	if err != nil {
		return err
	}

	fmt.Printf("Region: %s\n", region.Name)
	fmt.Printf("Main Generation: %s\n", region.MainGeneration.Name)
	fmt.Printf("Locations:\n")
	for _, location := range region.Locations {
		id := strings.Split(location.URL, "/")
		fmt.Printf(" %s. %s\n", id[len(id)-2], location.Name)
	}

	return nil
}
