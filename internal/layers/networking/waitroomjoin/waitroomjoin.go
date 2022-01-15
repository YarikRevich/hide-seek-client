package waitroomjoin

import (
	"context"
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_JOIN, func() {
		server := networking.UseNetworking().Clients().Base().GetClient()

		pcMess := world.UseWorld().GetPC().ToAPIMessage()

		if _, err := server.InsertOrUpdatePC(context.Background(), pcMess, grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		worldId := w.ID.String()

		server := networking.UseNetworking().Clients().Base().GetClient()

		worldObjects, err := server.FindWorldObjects(context.Background(), &wrappers.StringValue{Value: worldId}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}
		// fmt.Println(worldObjects)

		w.Update(worldObjects)

		fmt.Println(w.GetGameSettings().IsGameStarted)

		if !w.GetGameSettings().IsWorldExist {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})
			notifications.PopUp.WriteError("The room was closed by the creator")
		}
		if w.GetGameSettings().IsGameStarted {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})
		}
	}, statemachine.UI_WAIT_ROOM_JOIN, time.Second)
}
