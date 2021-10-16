package keyboard

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/game"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/settings_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
)

func Process() {
	switch statemachine.UseStateMachine().Input().GetState() {
	case input.SETTINGS_MENU_USERNAME:
		settingsmenu.Exec()
	case input.GAME:
		game.Exec()
	default:
		collection.CleanBlinking()
	}
}
