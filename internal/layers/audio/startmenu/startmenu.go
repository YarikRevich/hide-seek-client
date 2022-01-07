package startmenu

import "github.com/YarikRevich/hide-seek-client/internal/core/player"

func Exec() {
	player.UsePlayer().Play("background", player.PlayerOpts{Infinite: true})
}
