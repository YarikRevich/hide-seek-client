package waitroomstart

import (
	"context"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/atotto/clipboard"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() bool {
	m := events.UseEvents().Mouse()

	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			w := world.UseWorld()
			w.DeletePCs()

			conn := networking.UseNetworking().Dialer().Conn()
			if r, err := conn.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}); !r.Ok || err != nil {
				notifications.PopUp.WriteError("Can't delete the world")
				return true
			}

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_HERO_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("fonts/waitroom/waitroom").Modified) {
			if err := clipboard.WriteAll(world.UseWorld().String()); err != nil {
				logrus.Fatal(err)
			}
			notifications.PopUp.WriteError("World ID has been copied!")

			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified) {
			// worldId := objects.UseObjects().World().ID
			w := world.UseWorld()

			conn := networking.UseNetworking().Dialer().Conn()
			if r, err := conn.SetGameStarted(context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{}); !r.Ok || err != nil {
				notifications.PopUp.WriteError("Can't start game!")
				return true
			}

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_GAME)
			return true
		}
	}

	return false
}
