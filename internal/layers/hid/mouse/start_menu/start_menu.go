package startmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/settingswheel")) {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.SETTINGS_MENU),
				uimiddleware.UseUIMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_start")) {
			objects.UseObjects().PC().SetUsername()

			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.MAP_CHOOSE),
				uimiddleware.UseUIMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}

		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_join")) {
			objects.UseObjects().PC().SetUsername()

			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.JOIN_LOBBY_MENU),
				uimiddleware.UseUIMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}
	}
	return false
}
