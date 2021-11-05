package animation

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/animation/game"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.UI)
		defer profiling.UseProfiler().EndMonitoring()
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		func(){
			if cli.IsDebug(){
				profiling.UseProfiler().StartMonitoring(profiling.UI_GAME_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Exec()
		}()
	}
}
