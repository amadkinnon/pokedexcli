package cmds

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	area := args[0]
	fmt.Printf("Exploring %s...\n", area)

	pokemon, err := c.Client.GetLocation(area)
	if err != nil {
		fmt.Printf("Could not get any results for %s\n", area)
		return err
	}

	if len(pokemon.PokemonEncounters) == 0 {
		fmt.Printf("No pokemon found at %s\n", area)
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, p := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
