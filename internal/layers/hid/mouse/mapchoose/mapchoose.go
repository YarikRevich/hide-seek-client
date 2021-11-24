package mapchoose

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	m.UpdateMouseWheelOffsets()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
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
			if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata(k).Modified) {
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
