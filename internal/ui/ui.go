package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/debug"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/game"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/pop_up_messages"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/start_menu"
)

func Process() {
	switch ui.UseStatus().GetState() {
	case ui.GAME:
		game.Draw()
	case ui.START_MENU:
		startmenu.Draw()
	case ui.SETTINGS_MENU:
		settingsmenu.Draw()
	case ui.CREATE_LOBBY_MENU:
	case ui.JOIN_LOBBY_MENU:
		
	case ui.CHOOSE_EQUIPMENT:
	case ui.WAIT_ROOM:
	
	}
	debug.Draw()
	popupmessages.Draw()
}
