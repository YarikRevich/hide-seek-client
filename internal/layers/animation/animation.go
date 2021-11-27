package animation

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/animation/game"
	"github.com/YarikRevich/HideSeek-Client/tools/params"
)

func Process() {
	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		func() {
			if params.IsDebug() {
				profiling.UseProfiler().StartMonitoring(profiling.UI_ANIMATION)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Exec()
		}()
	}
}
