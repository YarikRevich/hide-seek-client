package waitroomstart

import (
	"context"
	"fmt"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_START, func() {
		conn := networking.UseNetworking().Dialer().Conn()

		w := world.UseWorld()
		worldMess := w.ToAPIMessage()
		mapMess := w.GetWorldMap().ToAPIMessage()
		pcMess := w.GetPC().ToAPIMessage()

		r, err := conn.AddWorld(context.Background(), worldMess, grpc.EmptyCallOption{})
		if !r.GetOk() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
		r, err = conn.AddMap(context.Background(), mapMess, grpc.EmptyCallOption{})
		if !r.GetOk() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
		r, err = conn.AddPC(context.Background(), pcMess, grpc.EmptyCallOption{})
		if !r.GetOk() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		conn := networking.UseNetworking().Dialer().Conn()

		fmt.Println("GAME STARTED BEFORE", w.GetGameSettings().IsGameStarted)
		r, err := conn.UpdateWorld(context.Background(), w.ToAPIMessage(), grpc.EmptyCallOption{})
		if !r.GetOk() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}

		if !w.GetGameSettings().IsWorldExist {
			r, err = conn.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
			if !r.GetOk() || err != nil {
				notifications.PopUp.WriteError(err.Error())
				return
			}
		}

		worldObjects, err := conn.GetWorld(
			context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		w.Update(worldObjects)

		fmt.Println("GAME STARTED AFTER", w.GetGameSettings().IsGameStarted)

	}, statemachine.UI_WAIT_ROOM_START, time.Second)
}
