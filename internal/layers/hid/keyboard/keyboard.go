package keyboard

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard/game"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard/settingsmenu"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard/joinmenu"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.KEYBOARD)
		defer profiling.UseProfiler().EndMonitoring()
	}

	switch statemachine.UseStateMachine().Input().GetState() {
	case statemachine.INPUT_SETTINGS_MENU_USERNAME:
		settingsmenu.Exec()
	case statemachine.INPUT_JOIN_MENU:
		joinmenu.Exec()
	case statemachine.INPUT_GAME:
		game.Exec()
	}
}
