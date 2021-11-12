package waitroomstart

import (
	"context"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	
	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			objects.UseObjects().World().ResetPCs()
			worldId := objects.UseObjects().World().ID
			if r, err := networking.UseNetworking().Dialer().Conn().RemoveWorld(context.Background(), &api.RemoveWorldRequest{WorldId: worldId.String()}); !r.Ok || err != nil{
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
			if err := clipboard.WriteAll(objects.UseObjects().World().ID.String()); err != nil{
				logrus.Fatal(err)
			}
			notifications.PopUp.WriteError("World ID has been copied!")

			return true
		}

		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified) {
			worldId := objects.UseObjects().World().ID
			if r, err := networking.UseNetworking().Dialer().Conn().SetGameStarted(context.Background(), &api.SetGameStartedRequest{WorldId: worldId.String()}, grpc.EmptyCallOption{}); !r.Ok || err != nil{
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
