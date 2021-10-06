package settingsmenu

import (

	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	// uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
)

func Exec()bool{
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/menues/inputs/input")) {
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Input().SetState(input.SETTINGS_MENU_USERNAME),
			inputmiddleware.UseInputMiddleware,
		)
		return true
	}
	return false
}