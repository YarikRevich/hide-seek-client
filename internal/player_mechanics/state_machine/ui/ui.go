package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
)

const (
	START_MENU statusEntry = iota
	SETTINGS_MENU

	CREATE_LOBBY_MENU
	JOIN_LOBBY_MENU

	CHOOSE_EQUIPMENT

	WAIT_ROOM

	GAME
)

var (
	stateMachine *Status
)

type statusEntry int

type Status struct {
	status statusEntry
}

func (s *Status) SetState(st statusEntry) {
	statemachine.UseMiddlewares()
	s.status = st
}

func (s *Status) GetState() statusEntry {
	return s.status
}

func UseStatus() *Status {
	if stateMachine == nil {
		stateMachine = &Status{status: START_MENU}
	}
	return stateMachine
}
