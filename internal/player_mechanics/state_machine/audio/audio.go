package audio

import (
	"sync"

	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/common"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
)

var instance common.IState

type Status struct {
	sync.Mutex
	status int
}

func (s *Status) SetState(st int) func() {
	return func() {
		s.Lock()
		defer s.Unlock()
		s.status = st
	}
}

func (s *Status) GetState() int {
	return s.status
}

func UseStatus() common.IState {
	if instance == nil {
		instance = &Status{status: audio.DONE}
	}
	return instance
}
