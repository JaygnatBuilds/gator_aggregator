package commands

import (
	state "github.com/JaygnatBuilds/gator_aggregator/internal/state"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state.State, cmd command) error {

}
