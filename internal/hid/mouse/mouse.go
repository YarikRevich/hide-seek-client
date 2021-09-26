package mouse

import (
	creationlobbymenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/creation_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/start_menu"
)

func Process() {
	switch statemachine.GetInstance().GetState(){
	case statemachine.CREATE_LOBBY_MENU:
		creationlobbymenu.Exec()
	case statemachine.START_MENU:
		startmenu.Exec()
	}
}