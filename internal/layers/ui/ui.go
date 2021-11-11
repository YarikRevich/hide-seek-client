package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/debug"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/game"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/herochoose"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/joinmenu"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/mapchoose"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/popup"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/settingsmenu"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/startmenu"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/waitroomjoin"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/waitroomstart"

	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.UI)
		defer profiling.UseProfiler().EndMonitoring()
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		func() {
			if cli.IsDebug() {
				profiling.UseProfiler().StartMonitoring(profiling.UI_GAME_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Draw()
		}()
	case statemachine.UI_START_MENU:
		func() {
			if cli.IsDebug() {
				profiling.UseProfiler().StartMonitoring(profiling.UI_START_MENU)
				defer profiling.UseProfiler().EndMonitoring()
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
