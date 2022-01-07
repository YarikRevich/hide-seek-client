package click

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
)

func Exec() {
	player.UsePlayer().Play("soundeffects/mousepress", player.PlayerOpts{})
}
