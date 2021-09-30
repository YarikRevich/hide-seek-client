package keyboard

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/game"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/start_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
)

func Process() {
	switch statemachine.GetInstance().GetState() {
	case statemachine.START_MENU:
		startmenu.Exec()
	case statemachine.SETTINGS_MENU:
	case statemachine.CREATE_LOBBY_MENU:
	case statemachine.JOIN_LOBBY_MENU:
	case statemachine.CHOOSE_EQUIPMENT:
	case statemachine.WAIT_ROOM:
	case statemachine.GAME:
		game.Exec()
	}
}
