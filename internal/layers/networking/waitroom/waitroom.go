package waitroom

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

func Exec() {
	//start:
	//add user to world

	//join:
	//get world
	//add user to world
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM, func() {

	})

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

	latency.UseLatency().Timings().ExecEach(func() {
		w := objects.UseObjects().World()
		networking.UseNetworking().Dialer().Conn().Call("update_world", w.ID, w)
	}, time.Second)
}
