package herochoose

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
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_MAP_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		for k, v := range map[string]string{
			"heroes/thumbnails/pumpkin": "heroes/pumpkin",
		} {
			if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata(k).Modified) {
				world.UseWorld().GetPC().SetSkin(v)

				switch statemachine.UseStateMachine().Game().GetState() {
				case statemachine.GAME_START:
					middlewares.UseMiddlewares().UI().UseAfter(func() {
						statemachine.UseStateMachine().UI().SetState(statemachine.UI_WAIT_ROOM_START)
					})
				case statemachine.GAME_JOIN:
					middlewares.UseMiddlewares().UI().UseAfter(func() {
						statemachine.UseStateMachine().UI().SetState(statemachine.UI_WAIT_ROOM_JOIN)
					})
				}

				statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
				return true
			}
		}
	}

	return false
}
