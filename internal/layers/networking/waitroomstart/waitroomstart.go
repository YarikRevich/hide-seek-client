package waitroomstart

import (
	"context"
	"fmt"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_START, func() {
		fmt.Println("HERE")

		conn := networking.UseNetworking().Dialer().Conn()

		w := world.UseWorld()
		worldMess := w.ToAPIMessage()
		mapMess := w.GetWorldMap().ToAPIMessage()
		pcMess := w.GetPC().ToAPIMessage()

		conn.AddWorld(context.Background(), worldMess, grpc.EmptyCallOption{})
		conn.AddMap(context.Background(), mapMess, grpc.EmptyCallOption{})
		conn.AddPC(context.Background(), pcMess, grpc.EmptyCallOption{})
	})

	latency.UseLatency().Timings().ExecEach(func() {
		fmt.Println("HERE1")

		w := world.UseWorld()

		conn := networking.UseNetworking().Dialer().Conn()
		worldObjects, err := conn.GetWorldProperty(
			context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		w.UpdateProperty(worldObjects)
	}, statemachine.UI_WAIT_ROOM_START, time.Second)
}
