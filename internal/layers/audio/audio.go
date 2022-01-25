package audio

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/buttonclick"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/click"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/game"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/notificationnew"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/startmenu"
	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.AUDIO)
		defer ingame.UseProfiler().StopMonitoring(ingame.AUDIO)
	}

	switch statemachine.UseStateMachine().Mouse().GetState() {
	case statemachine.MOUSE_BUTTON_CLICK:
		buttonclick.Exec()
	case statemachine.MOUSE_CLICK:
		click.Exec()
	}
	statemachine.UseStateMachine().Mouse().SetState(statemachine.MOUSE_NONE)

	switch statemachine.UseStateMachine().Notification().GetState() {
	case statemachine.NOTIFICATION_NEW:
		fmt.Println("HERE")
		notificationnew.Exec()
	}
	statemachine.UseStateMachine().Notification().SetState(statemachine.NOTIFICATION_NONE)

	if statemachine.UseStateMachine().Audio().GetState() == statemachine.AUDIO_DONE {
		switch statemachine.UseStateMachine().UI().GetState() {
		case statemachine.UI_START_MENU:
			startmenu.Exec()
		case statemachine.UI_GAME:
			game.Exec()
		default:
			return
		}

		statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_UNDONE)
	}

}
