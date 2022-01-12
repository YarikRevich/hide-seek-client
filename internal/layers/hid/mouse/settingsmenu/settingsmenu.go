package settingsmenu

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/storage"
	"github.com/YarikRevich/hide-seek-client/tools/lanserver"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	if m.IsAnyMouseButtonsPressed() {

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/inputs/settingsmenuinput")) {

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_SETTINGS_MENU_USERNAME)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back")) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config")) {
			storage.UseStorage().User().SetUsername(events.UseEvents().Input().SettingsMenuNameBuffer.ReadClean())

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxoff")) ||
			m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxon")) {
			switch statemachine.UseStateMachine().SettingsMenuCheckbox().GetState() {
			case statemachine.UI_SETTINGS_MENU_CHECKBOX_OFF:
				statemachine.UseStateMachine().SettingsMenuCheckbox().SetState(statemachine.UI_SETTINGS_MENU_CHECKBOX_ON)
				statemachine.UseStateMachine().Dial().SetState(statemachine.DIAL_LAN)
				lanserver.UseLANServer().Start()
			case statemachine.UI_SETTINGS_MENU_CHECKBOX_ON:
				statemachine.UseStateMachine().SettingsMenuCheckbox().SetState(statemachine.UI_SETTINGS_MENU_CHECKBOX_OFF)
				statemachine.UseStateMachine().Dial().SetState(statemachine.DIAL_WAN)
				lanserver.UseLANServer().Stop()
			}

			return true
		}
	}

	return false
}
