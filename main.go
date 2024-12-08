package main

import (
	"pokedexcli/internal/api"
	"pokedexcli/internal/cmds"
	"pokedexcli/internal/repl"
	"time"
)

func main() {
	apiclient := api.NewClient(5*time.Second, 5*time.Minute)
	cfg := &cmds.Config{
		Caught: map[string]api.Pokemon{},
		Client: apiclient,
	}

	repl.StartRepl(cfg)
}
