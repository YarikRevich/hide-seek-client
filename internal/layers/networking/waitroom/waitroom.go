package waitroom

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/timings"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
)

func Exec() {
	//start:
	//add user to world

	//join:
	//get world
	//add user to world
	collection.OnceCollection[ui.WAIT_ROOM].Do(func() {
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
	})

	timings.UseTimings().ExecEach(func() {
		w := objects.UseObjects().World()
		connection.UseConnection().Call("update_world", w.ID, w)
	}, time.Second)
}
