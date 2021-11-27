package audio

import (
	// "github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/audio/game"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/audio/startmenu"

	"github.com/YarikRevich/HideSeek-Client/tools/params"
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

		// middlewares.UseMiddlewares().Audio().UseAfter(func() {
		statemachine.UseStateMachine().Audio().SetState(statemachine.AUDIO_UNDONE)
		// })
	}
}
