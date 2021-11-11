package joinmenu


import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/storage"
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
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_join_world").Modified) {
			storage.UseStorage().User().SetUsername(events.UseEvents().Input().JoinGameBuffer.ReadClean())

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
	}

	return false
}
