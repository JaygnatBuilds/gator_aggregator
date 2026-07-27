package commands

import (
	"fmt"

	state "github.com/JaygnatBuilds/gator_aggregator/internal/state"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state.State, cmd command) error {

	// if command argument slice is empty return an error
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login command expects username parameter.")
	}

	// retrieve username from command args
	user := cmd.args[0]

	// update username field in config struct
	s.Config.DB_USR = user

	return nil

}
