package waitroom

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
)

var ticker = time.NewTicker(time.Second)

func Exec() {
	collection.OnceCollection[ui.WAIT_ROOM].Do(func() {
		connection.UseConnection().Call("reg_user", pc.UsePC(), nil)

		connection.UseConnection().Call("reg_world", struct {
			World world.World
			PC    pc.PC
		}{
			*world.UseWorld(), *pc.UsePC(),
		}, nil)
	})

	select {
	case <-ticker.C:
		connection.UseConnection().Call("update_world_users", world.UseWorld(), &world.UseWorld().Users)
	default:
	}
}
