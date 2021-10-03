package keyboard

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/game"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/start_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/input"
)

func Process() {
	switch input.UseStatus().GetState() {
	case input.SETTINGS_MENU_USERNAME:
		startmenu.Exec()
	case input.GAME:
		game.Exec()
	}
}
