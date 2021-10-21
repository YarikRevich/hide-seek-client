package statemachine

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/common"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"
)

var instance IStateMachine

type StateMachine struct{}

type IStateMachine interface {
	Audio() common.IState
	UI() common.IState
	Input() common.IState
	Networking() common.IState
}

func (s *StateMachine) Audio() common.IState {
	return audio.UseStatus()
}

func (s *StateMachine) UI() common.IState {
	return ui.UseStatus()
}

func (s *StateMachine) Input() common.IState {
	return input.UseStatus()
}

func (s *StateMachine) Networking() common.IState {
	return networking.UseStatus()
}

func UseStateMachine() IStateMachine {
	if instance == nil {
		instance = new(StateMachine)
	}
	return instance
}
