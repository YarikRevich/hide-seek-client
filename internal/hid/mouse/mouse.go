package mouse

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/join_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/unfocus"
	waitroom "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/wait_room"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"

	settingsmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/settings_menu"
	startmenu "github.com/YarikRevich/HideSeek-Client/internal/hid/mouse/start_menu"
)

func Process() {
	profiling.UseProfiler().StartMonitoring(profiling.MOUSE_HANDLER)

	if statemachine.UseStateMachine().Networking().GetState() == networking.ONLINE {
		switch statemachine.UseStateMachine().UI().GetState() {
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
		}

		if IsMousePressed() {
			unfocus.Exec()
		}
	}

	profiling.UseProfiler().EndMonitoring()
}
