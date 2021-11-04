package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/debug"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/game"
	herochoose "github.com/YarikRevich/HideSeek-Client/internal/layers/ui/hero_choose"
	mapchoose "github.com/YarikRevich/HideSeek-Client/internal/layers/ui/map_choose"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui/popup"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/layers/ui/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/layers/ui/start_menu"
	waitroom "github.com/YarikRevich/HideSeek-Client/internal/layers/ui/wait_room"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.UI)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
			debug.Draw()
		}()
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case ui.GAME:
		func(){
			if cli.IsDebug(){
				profiling.UseProfiler().StartMonitoring(profiling.UI_GAME_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Draw()
		}()
	case ui.START_MENU:
		func(){
			if cli.IsDebug(){
				profiling.UseProfiler().StartMonitoring(profiling.UI_START_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			startmenu.Draw()
		}()
	case ui.SETTINGS_MENU:
		settingsmenu.Draw()
	case ui.MAP_CHOOSE:
		mapchoose.Draw()
	case ui.HERO_CHOOSE:
		herochoose.Draw()
	case ui.JOIN_LOBBY_MENU:

	case ui.CHOOSE_EQUIPMENT:
	case ui.WAIT_ROOM:
		waitroom.Draw()
	}
	popup.Draw()
}
