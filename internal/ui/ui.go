package ui

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/debug"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/game"
	herochoose "github.com/YarikRevich/HideSeek-Client/internal/ui/hero_choose"
	mapchoose "github.com/YarikRevich/HideSeek-Client/internal/ui/map_choose"
	"github.com/YarikRevich/HideSeek-Client/internal/ui/pop_up_messages"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/ui/start_menu"
	waitroom "github.com/YarikRevich/HideSeek-Client/internal/ui/wait_room"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.GetDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.UI)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
			debug.Draw()
		}()
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case ui.GAME:
		func(){
			if cli.GetDebug(){
				profiling.UseProfiler().StartMonitoring(profiling.UI_GAME_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Draw()
		}()
	case ui.START_MENU:
		func(){
			if cli.GetDebug(){
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
	popupmessages.Draw()
}
