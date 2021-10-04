package keyboard

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/game"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/start_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
)

func Process() {
	switch statemachine.UseStateMachine().Input().GetState() {
	case input.SETTINGS_MENU_USERNAME:
		startmenu.Exec()
	case input.GAME:
		game.Exec()
	}
}
