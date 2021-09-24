package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/pop_up_messages"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/start_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/debug"
)

func Process() {
	switch statemachine.GetInstance().GetState() {
	case statemachine.START_MENU:
		start_menu.Draw()
	case statemachine.SETTINGS_MENU:
	case statemachine.CREATE_LOBBY_MENU:
	case statemachine.JOIN_LOBBY_MENU:
		
	case statemachine.CHOOSE_EQUIPMENT:
	case statemachine.WAIT_ROOM:
	case statemachine.GAME:

	}
	debug.Draw()
	popupmessages.Draw()
}
