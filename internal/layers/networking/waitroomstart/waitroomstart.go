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

		r, err := server.UpdateWorld(context.Background(), worldMess, grpc.EmptyCallOption{})
		if !r.GetValue() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
		r, err = server.UpdateMap(context.Background(), mapMess, grpc.EmptyCallOption{})
		if !r.GetValue() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
		r, err = server.UpdatePC(context.Background(), pcMess, grpc.EmptyCallOption{})
		if !r.GetValue() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		server := networking.UseNetworking().Clients().Base().GetClient()

		fmt.Println("GAME STARTED BEFORE", w.GetGameSettings().IsGameStarted)
		r, err := server.UpdateWorld(context.Background(), w.ToAPIMessage(), grpc.EmptyCallOption{})
		if !r.GetValue() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}

		if !w.GetGameSettings().IsWorldExist {
			r, err = server.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
			if !r.GetValue() || err != nil {
				notifications.PopUp.WriteError(err.Error())
				return
			}
		}

		worldObjects, err := server.GetWorld(
			context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		w.Update(worldObjects)

		fmt.Println("GAME STARTED AFTER", w.GetGameSettings().IsGameStarted)

	}, statemachine.UI_WAIT_ROOM_START, time.Second)
}
