package animation

import (

	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/animation/game"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
)

func Process() {
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.UI)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
		}()
	}

	switch statemachine.UseStateMachine().UI().GetState() {
	case ui.GAME:
		func(){
			if cli.IsDebug(){
				profiling.UseProfiler().StartMonitoring(profiling.UI_GAME_MENU)
				defer profiling.UseProfiler().EndMonitoring()
			}
			game.Exec()
		}()
	}
}
