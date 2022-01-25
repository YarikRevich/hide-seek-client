package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/debug"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/game"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/herochoose"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/joinmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/mapchoose"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/popup"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/settingsmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/startmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/waitroomjoin"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui/waitroomstart"

	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.UI)
		defer ingame.UseProfiler().StopMonitoring(ingame.UI)
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		func() {
			if params.IsDebug() {
				ingame.UseProfiler().StartMonitoring(ingame.UI_GAME_MENU)
				defer ingame.UseProfiler().StopMonitoring(ingame.UI_GAME_MENU)
			}
			game.Draw()
		}()
	case statemachine.UI_START_MENU:
		func() {
			if params.IsDebug() {
				ingame.UseProfiler().StartMonitoring(ingame.UI_START_MENU)
				defer ingame.UseProfiler().StopMonitoring(ingame.UI_START_MENU)
			}
			startmenu.Draw()
		}()
	case statemachine.UI_SETTINGS_MENU:
		settingsmenu.Draw()
	case statemachine.UI_MAP_CHOOSE:
		mapchoose.Draw()
	case statemachine.UI_HERO_CHOOSE:
		herochoose.Draw()
	case statemachine.UI_JOIN_MENU:
		joinmenu.Draw()
	case statemachine.UI_WAIT_ROOM_START:
		waitroomstart.Draw()
	case statemachine.UI_WAIT_ROOM_JOIN:
		waitroomjoin.Draw()
	case statemachine.UI_CHOOSE_EQUIPMENT:

	}
	popup.Draw()

	debug.Draw()
}
