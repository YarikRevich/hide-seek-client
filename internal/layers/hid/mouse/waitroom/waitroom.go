package waitroom

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	
	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			objects.UseObjects().World().ResetPCs()

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_MAP_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("fonts/waitroom/waitroom").Modified) {
			if err := clipboard.WriteAll(objects.UseObjects().World().ID.String()); err != nil{
				logrus.Fatal(err)
			}
			notifications.PopUp.WriteError("World ID has been copied!")

			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified) {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_GAME)
			return true
		}
	}

	return false
}
