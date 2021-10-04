package input

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/common"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
)

var instance *Status

type Status struct {
	status int
}

func (s *Status) SetState(st int) func() {
	return func() {
		s.status = st
	}
}

func (s *Status) GetState() int {
	return s.status
}

func UseStatus() common.IState {
	if instance == nil {
		instance = &Status{status: input.EMPTY}
	}
	return instance
}
