package waitroomstart

import (
	"context"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

func Exec() {
	//start:
	//add user to world

	//join:
	//get world
	//add user to world
	// latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM, func() {
	// 	networking.UseNetworking().Dialer().Conn().Call("add_user_to_world", w, )
	// })

	// collection.OnceCollection[ui.WAIT_ROOM].Do(func() {
	// o := objects.UseObjects()
	// connection.UseConnection().Call("reg_user", o.PC(), nil)

	// connection.UseConnection().Call("reg_world", struct {
	// 	World objects.World
	// 	PC    objects.PC
	// }{
	// 	*o.World(), *o.PC(),
	// }, nil)
	// struct{}{}
	// connection.UseConnection().Call("add_user_to_world", nil, nil)
	// })
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_START, func() {
		// m := objects.UseObjects().World()
		// networking.UseNetworking().Dialer().Conn().AddWorld(context.Background(), m, grpc.EmptyCallOption{})
		// wm := objects.UseObjects().World().ToAPIMessage()
		// networking.UseNetworking().Dialer().Conn().AddWorld(context.Background(), wm, grpc.EmptyCallOption{})

		// pm := objects.UseObjects().PC().ToAPIMessage()
		// networking.UseNetworking().Dialer().Conn().AddPC(context.Background(), pm, grpc.EmptyCallOption{})

	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		conn := networking.UseNetworking().Dialer().Conn()
		worldObjects, err := conn.GetWorldObjects(
			context.Background(), &api.WorldObjectsRequest{WorldId: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		w.UpdateProperty(worldObjects)
	}, statemachine.UI_WAIT_ROOM_START, time.Second)
}
