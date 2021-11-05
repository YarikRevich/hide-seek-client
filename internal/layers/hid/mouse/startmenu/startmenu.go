package startmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/settingswheel").Modified) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_SETTINGS_MENU)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(		sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/button_start").Modified) {
			objects.UseObjects().PC().SetUsername()

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_MAP_CHOOSE)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/button_join").Modified) {
			objects.UseObjects().PC().SetUsername()

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_JOIN_LOBBY_MENU)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
	}
	return false
}
