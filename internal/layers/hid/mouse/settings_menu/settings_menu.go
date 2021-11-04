package settingsmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/provider"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/inputs/input")) {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.SETTINGS_MENU_USERNAME),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.START_MENU),
				uimiddleware.UseUIMiddleware,
			)
			return true
		}
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_save_config")) {
			provider.UseStorageProvider().SetUsername(events.UseEvents().Input().SettingsMenuNameBuffer.ReadClean())

			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.START_MENU),
				uimiddleware.UseUIMiddleware,
			)
			return true
		}
	}

	return false
}
