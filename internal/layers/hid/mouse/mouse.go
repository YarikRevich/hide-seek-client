package mouse

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/herochoose"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/joinmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/mapchoose"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/settingsmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/startmenu"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/unfocus"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/waitroomjoin"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse/waitroomstart"

	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.MOUSE)
		defer ingame.UseProfiler().StopMonitoring(ingame.MOUSE)
	}

	if statemachine.UseStateMachine().Networking().GetState() == statemachine.NETWORKING_ONLINE {
		switch statemachine.UseStateMachine().UI().GetState() {
		case statemachine.UI_GAME:
			return
		case statemachine.UI_JOIN_MENU:
			if joinmenu.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_WAIT_ROOM_START:
			if waitroomstart.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_WAIT_ROOM_JOIN:
			if waitroomjoin.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_START_MENU:
			if startmenu.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_SETTINGS_MENU:
			if settingsmenu.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_MAP_CHOOSE:
			if mapchoose.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		case statemachine.UI_HERO_CHOOSE:
			if herochoose.Exec() {
				statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_BUTTON_CLICK)
				return
			}
		}

		if events.UseEvents().Mouse().IsAnyMouseButtonsPressed() {
			statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_CLICK)
			unfocus.Exec()
		}
	}
}
