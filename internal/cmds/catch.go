package cmds

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")

	}

	name := args[0]
	pokemon, err := c.Client.GetPokemon(name)
	if err != nil {
		fmt.Printf("Could not get any details for %s\n", name)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s(%d)...\n", pokemon.Name, pokemon.BaseExperience)
	res := rand.Intn(pokemon.BaseExperience)

	if res > 40 {
		fmt.Printf("%s was caught!\n\n", name)
		return nil
	}
	fmt.Printf("%s escaped!\n\n", name)
	return nil
}
