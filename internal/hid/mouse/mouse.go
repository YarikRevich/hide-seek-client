package mouse

import (
	creationlobbymenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/creation_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
)

func Process() {
	switch statemachine.GetInstance().GetState(){
	case statemachine.CREATE_LOBBY_MENU:
		creationlobbymenu.Exec()
	}
}