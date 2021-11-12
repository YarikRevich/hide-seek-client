package joinmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/google/uuid"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	if m.IsAnyMouseButtonsPressed() {

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/inputs/joingameinput").Modified) {
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_JOIN_MENU)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_join_game").Modified) {
			worldId := events.UseEvents().Input().JoinGameBuffer.ReadClean()
			id, err := uuid.Parse(worldId)
			if err != nil {
				notifications.PopUp.WriteError("ID is malformed!")
				return true
			}
			objects.UseObjects().World().SetID(id)

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_HERO_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
	}

	return false
}
