package main

import (
	"fmt"
	"strings"
)

func runLocation(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Location ID or name")
		return nil
	}

	idOrName := strings.Join(args, " ")

	area, err := conf.pokeapiClient.GetLocation(idOrName)

	if err != nil {
		return err
	}

	regionName := area.Name
	if area.Region != nil {
		regionName = area.Region.Name
	}

	fmt.Printf("Location: %s\n", area.Name)
	fmt.Printf("Region: %s\n", regionName)

	encounters := area.Areas

	if len(encounters) == 0 {
		fmt.Println("Areas: None")
		return nil
	}

	fmt.Printf("Areas:\n")
	for _, encounter := range encounters {
		id := strings.Split(encounter.URL, "/")
		fmt.Printf("  %s. %s\n", id[len(id)-2], encounter.Name)
	}

	return nil
}
