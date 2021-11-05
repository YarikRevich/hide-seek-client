package herochoose

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
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_MAP_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		for k, v := range map[string]string{
			"assets/images/heroes/thumbnails/pumpkin": "assets/images/heroes/pumpkin",
		} {
			if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata(k).Modified) {
				p := objects.UseObjects().PC()
				p.SetSkin(v)
				objects.UseObjects().World().AddPC(p)

				middlewares.UseMiddlewares().UI().UseAfter(func() {
					statemachine.UseStateMachine().UI().SetState(statemachine.UI_WAIT_ROOM)
				})

				statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
				return true
			}
		}
	}

	return false
}