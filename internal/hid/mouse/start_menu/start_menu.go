package startmenu

import (
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/world"
)

func Exec()bool {
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/settingswheel")) {
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
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_start")) {
		pc.UsePC().Init()
		world.UseWorld().Init("assets/images/maps/helloween/background/background")
		
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().UI().SetState(ui.WAIT_ROOM),
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
