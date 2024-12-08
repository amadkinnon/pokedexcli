package cmds

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config, args ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a pokemon",
			Callback:    commandCatch,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a given area",
			Callback:    commandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Displays the name of 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map",
			Description: "Displays the name of previous 20 locations",
			Callback:    commandMapb,
		},
	}
}
