package game

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

func Show() {
	world.UseWorld().DebugInit()

	middlewares.UseMiddlewares().UI().UseAfter(func() {
		statemachine.Layers.SetState(statemachine.LAYERS_SESSION)
	})
	statemachine.Input.SetState(statemachine.INPUT_GAME)
}
