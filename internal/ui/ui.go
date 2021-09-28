package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/debug"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/game"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/pop_up_messages"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/start_menu"
)

func Process() {
	switch statemachine.GetInstance().GetState() {
	case statemachine.GAME:
		game.Draw()
	case statemachine.START_MENU:
		startmenu.Draw()
	case statemachine.SETTINGS_MENU:
		settingsmenu.Draw()
	case statemachine.CREATE_LOBBY_MENU:
	case statemachine.JOIN_LOBBY_MENU:
		
	case statemachine.CHOOSE_EQUIPMENT:
	case statemachine.WAIT_ROOM:
	
	}
	debug.Draw()
	popupmessages.Draw()
}
