package cmds

import "fmt"

func commandHelp(c *Config, arg ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for name, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}
