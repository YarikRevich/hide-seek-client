package waitroom

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
)

var ticker = time.NewTicker(time.Second)

func Exec() {
	collection.OnceCollection[ui.WAIT_ROOM].Do(func() {
		o := objects.UseObjects()
		connection.UseConnection().Call("reg_user", o.PC(), nil)

		connection.UseConnection().Call("reg_world", struct {
			World objects.World
			PC    objects.PC
		}{
			*o.World(), *o.PC(),
		}, nil)
	})

	select {
	case <-ticker.C:
		w := objects.UseObjects().World()
		connection.UseConnection().Call("update_world_users", w, &w.PCs)
	default:
	}
}
