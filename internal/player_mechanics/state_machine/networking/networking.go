package networking

import (
	"sync"

	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/common"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
)

var instance common.IState

type Status struct {
	sync.Mutex
	status int
}

func (s *Status) SetState(st int) func() {
	return func() {
		s.status = st
	}
}

func (s *Status) GetState() int {
	s.Lock()
	defer s.Unlock()
	return s.status
}

func UseStatus() common.IState {
	if instance == nil {
		instance = &Status{status: networking.OFFLINE}
	}
	return instance
}
