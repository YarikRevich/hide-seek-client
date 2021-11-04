package mouse

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/join_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/unfocus"
	waitroom "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/wait_room"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"

	herochoose "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/hero_choose"
	mapchoose "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/map_choose"
	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/start_menu"
)

func Process() {
	if cli.GetDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.MOUSE)
		defer profiling.UseProfiler().EndMonitoring()
	}

	if statemachine.UseStateMachine().Networking().GetState() == networking.ONLINE {
		switch statemachine.UseStateMachine().UI().GetState() {
		case ui.GAME:
			return
		case ui.WAIT_ROOM:
			if waitroom.Exec() {
				return
			}
		case ui.JOIN_LOBBY_MENU:
			if joinlobbymenu.Exec() {
				return
			}
		case ui.START_MENU:
			if startmenu.Exec() {
				return
			}

		case ui.SETTINGS_MENU:
			if settingsmenu.Exec() {
				return
			}
		case ui.MAP_CHOOSE:
			if mapchoose.Exec() {
				return
			}
		case ui.HERO_CHOOSE:
			if herochoose.Exec() {
				return
			}
		}

		if events.UseEvents().Mouse().IsAnyMouseButtonsPressed() {
			unfocus.Exec()
		}
	}
}
