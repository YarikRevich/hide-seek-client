package waitroom

import (


	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
			objects.UseObjects().World().ResetPCs()

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

		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/fonts/waitroom/waitroom")) {
			if err := clipboard.WriteAll(objects.UseObjects().World().ID.String()); err != nil{
				logrus.Fatal(err)
			}
			notifications.PopUp.WriteError("World ID has been copied!")

			return true
		}

		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/button_confirm_game")) {

			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.GAME),
				uimiddleware.UseUIMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.GAME),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}
	}

	return false
}
