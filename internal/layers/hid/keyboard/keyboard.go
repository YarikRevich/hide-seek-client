package keyboard

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/keyboard/game"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/keyboard/joinmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/keyboard/settingsmenu"
	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	if params.IsDebug() {
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
