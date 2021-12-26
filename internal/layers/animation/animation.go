package animation

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/YarikRevich/hide-seek-client/internal/layers/animation/game"
	"github.com/YarikRevich/hide-seek-client/tools/params"
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
