package mapchoose

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	m.UpdateMouseWheelOffsets()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/back").Modified) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		for k, v := range map[string]string{
			"assets/images/maps/thumbnails/helloween": "assets/images/maps/helloween/background/background",
			"assets/images/maps/thumbnails/starwars":  "assets/images/maps/starwars/background/background",
		} {
			if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata(k).Modified) {
				objects.UseObjects().World().SetSkin(v)

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
