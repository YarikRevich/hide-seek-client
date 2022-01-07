package audio

import (
	// "github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/game"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio/startmenu"

	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	if params.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.AUDIO)
		defer profiling.UseProfiler().EndMonitoring()
	}

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

	switch statemachine.UseStateMachine().Mouse().GetState() {
	case statemachine.MOUSE_BUTTON_CLICK:
		// startmenu.Exec()
	case statemachine.MOUSE_CLICK:
		// game.Exec()
	default:
		return
	}
}
