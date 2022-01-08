package waitroomstart

import (
	"context"
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_START, func() {
		server := networking.UseNetworking().Clients().Base().GetClient()

		w := world.UseWorld()
		worldMess := w.ToAPIMessage()
		mapMess := w.GetWorldMap().ToAPIMessage()
		pcMess := w.GetPC().ToAPIMessage()
		// fmt.Println("HERE1")
		if _, err := server.InsertOrUpdateWorld(context.Background(), worldMess, grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			logrus.Error(err)
			return
		}
		// fmt.Println("HERE2", mapMess)
		if _, err := server.InsertOrUpdateMap(context.Background(), mapMess, grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			logrus.Error(err)
			return
		}
		// fmt.Println("HERE3", pcMess)
		if _, err := server.InsertOrUpdatePC(context.Background(), pcMess, grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			logrus.Error(err)
			return
		}

	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		server := networking.UseNetworking().Clients().Base().GetClient()

		// fmt.Println("GAME STARTED BEFORE", w.GetGameSettings().IsGameStarted)
		if _, err := server.InsertOrUpdateWorld(context.Background(), w.ToAPIMessage(), grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			logrus.Error(err)
			return
		}

		// if !w.GetGameSettings().IsWorldExist {
		// 	if _, err := server.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{}); err != nil {
		// 		notifications.PopUp.WriteError(err.Error())
		// 		logrus.Error(err)
		// 		return
		// 	}
		// }

		worldObjects, err := server.FindWorldObjects(
			context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			notifications.PopUp.WriteError(err.Error())
			logrus.Fatal(err)
		}
		// fmt.Println(worldObjects)

		w.Update(worldObjects)

		fmt.Println("HERER")

		// fmt.Println("GAME STARTED AFTER", w.GetGameSettings().IsGameStarted)

	}, statemachine.UI_WAIT_ROOM_START, time.Second)
}
