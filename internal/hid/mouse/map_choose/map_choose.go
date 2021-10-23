package map_choose

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
)

func Exec()bool {
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().UI().SetState(ui.START_MENU),
			uimiddleware.UseUIMiddleware,
		)
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Input().SetState(input.EMPTY),
			inputmiddleware.UseInputMiddleware,
		)
		return true
	}
	return false
}
