package waitroomstart

import (
	// "context"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"

	// "github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/atotto/clipboard"

	// "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			// w := world.UseWorld()
			// w.DeletePCs()
			world.UseWorld().GetGameSettings().SetWorldExist(false)

			// conn := networking.UseNetworking().Dialer().Conn()
			// if r, err := conn.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}); !r.Ok || err != nil {
			// 	notifications.PopUp.WriteError("Can't delete the world")
			// 	return true
			// }

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_HERO_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("fonts/waitroom/waitroom").Modified) {
			if err := clipboard.WriteAll(world.UseWorld().ID.String()); err != nil {
				logrus.Fatal(err)
			}
			notifications.PopUp.WriteError("World ID has been copied!")

			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified) {
			world.UseWorld().GetGameSettings().SetGameStarted(true)

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_GAME)
			return true
		}
	}

	return false
}
