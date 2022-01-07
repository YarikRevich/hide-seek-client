package buttonclick

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
)

func Exec() {
	player.UsePlayer().Play("soundeffects/mousepressbutton", player.PlayerOpts{})
}
