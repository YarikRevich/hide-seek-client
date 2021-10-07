package mouse

import (
	creationlobbymenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/creation_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/unfocus"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"

	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/start_menu"
)

func Process() {
	switch statemachine.UseStateMachine().UI().GetState(){
	case ui.CREATE_LOBBY_MENU:
		if creationlobbymenu.Exec(){
			return
		}
	case ui.START_MENU:
		if startmenu.Exec(){
			return 
		}
	case ui.SETTINGS_MENU:
		if settingsmenu.Exec(){
			return
		}
	}

	if IsMousePressed(){
		unfocus.Exec()
	}
}