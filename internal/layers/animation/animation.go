package animation

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/YarikRevich/hide-seek-client/internal/layers/animation/game"
	"github.com/YarikRevich/hide-seek-client/tools/params"
)

func Process() {
	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		func() {
			if params.IsDebug() {
				ingame.UseProfiler().StartMonitoring(ingame.UI_ANIMATION)
				defer ingame.UseProfiler().StopMonitoring(ingame.UI_ANIMATION)
			}
			game.Exec()
		}()
	}
}
