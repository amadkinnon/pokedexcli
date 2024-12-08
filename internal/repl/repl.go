package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/cmds"
	"strings"
)

const PROMPT = "pokedex > "

func StartRepl(config *cmds.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)

		scanner.Scan()
		input := scanner.Text()
		words := strings.Split(input, " ")
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := cmds.GetCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := cmd.Callback(config, args...)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}
