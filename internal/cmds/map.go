package cmds

import (
	"fmt"
)

func commandMap(c *Config, args ...string) error {
	locations, err := c.Client.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Prev = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}
	return nil
}

func commandMapb(c *Config, arg ...string) error {
	if c.Prev == nil {
		fmt.Println("At the first page of areas")
		return nil
	}

	locations, err := c.Client.ListLocations(c.Prev)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Prev = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}
	return nil
}
