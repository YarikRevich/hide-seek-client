package mapchoose

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	m.UpdateMouseWheelOffsets()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back")) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		for k, v := range map[string]string{
			"maps/thumbnails/helloween": "maps/helloween/background/background",
			"maps/thumbnails/starwars":  "maps/starwars/background/background",
		} {
			if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata(k)) {
				world.UseWorld().GetWorldMap().SetSkin(v)

				middlewares.UseMiddlewares().UI().UseAfter(func() {
					statemachine.UseStateMachine().UI().SetState(statemachine.UI_HERO_CHOOSE)
				})

				statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
				return true
			}
		}
	}

	return false
}
