package waitroom

import (
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func Exec() bool {
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
		world.UseWorld().ResetUsers()

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
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_confirm_game")) {
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Input().SetState(input.GAME),
			inputmiddleware.UseInputMiddleware,
		)
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().UI().SetState(ui.GAME),
			uimiddleware.UseUIMiddleware,
		)
		return true
	}
	return false
}
