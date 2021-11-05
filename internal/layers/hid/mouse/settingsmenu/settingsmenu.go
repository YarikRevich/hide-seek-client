package settingsmenu

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
		
		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/inputs/input").Modified) {
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_SETTINGS_MENU_USERNAME)
			return true
		}
		
		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/back").Modified) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
		if m.IsMousePressLeftOnce(sources.UseSources().Metadata().GetMetadata("assets/images/system/buttons/button_save_config").Modified) {
			storage.UseStorage().User().SetUsername(events.UseEvents().Input().SettingsMenuNameBuffer.ReadClean())

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
	}

	return false
}