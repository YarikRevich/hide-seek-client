package herochoose

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
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_MAP_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		for k, v := range map[string]string{
			"heroes/thumbnails/pumpkin": "heroes/pumpkin",
		} {
			if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata(k)) {
				w := world.UseWorld()
				w.GetPC().SetSkin(v)
				w.GetGameSettings().SetWorldExist(true)

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
