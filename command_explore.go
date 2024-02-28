package main

import (
	"fmt"
	"strings"
)

func runExplore(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Area ID or name")
		return nil
	}

	idOrName := strings.Join(args, " ")

	area, err := conf.pokeapiClient.GetLocationArea(idOrName)

	if err != nil {
		return err
	}

	cleanedAreaName := strings.Replace(area.Name, fmt.Sprintf("%s-", area.Location.Name), "", -1)
	if cleanedAreaName == "area" {
		cleanedAreaName = fmt.Sprintf("Common - (%s)", area.Name)
	}
	fmt.Printf("Area: %s\nLocation: %s\n", cleanedAreaName, area.Location.Name)

	encounters := area.PokemonEncounters

	if len(encounters) == 0 {
		fmt.Println("Encounters: None found.")
		return nil
	}

	fmt.Printf("Encounters:\n")
	for _, encounter := range encounters {
		fmt.Printf(" - %s: %d%%\n", encounter.Pokemon.Name, encounter.VersionDetails[0].EncounterDetails[0].Chance)
	}

	return nil
}
