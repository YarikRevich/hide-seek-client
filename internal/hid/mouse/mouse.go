package mouse

import (
	creationlobbymenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/creation_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/start_menu"
)

func Process() {
	switch ui.UseStatus().GetState(){
	case ui.CREATE_LOBBY_MENU:
		creationlobbymenu.Exec()
	case ui.START_MENU:
		startmenu.Exec()
	}
}