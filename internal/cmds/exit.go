package cmds

import "os"

func commandExit(c *Config, arg ...string) error {
	os.Exit(0)
	return nil
}
