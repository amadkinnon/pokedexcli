package cmds

import (
	"pokedexcli/internal/api"
)

type Config struct {
	Client api.Client
	Next   *string
	Prev   *string
	Caught map[string]api.Pokemon
}
