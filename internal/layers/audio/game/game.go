package game

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

func Exec() {
	p := player.UsePlayer()
	worldMap := world.UseWorld().GetWorldMap()
	switch worldMap.Name {
	case "helloween":
		p.Play("game", player.PlayerOpts{Infinite: true})
	case "starwars":
		p.Play("starwarsmaptheme", player.PlayerOpts{Infinite: true})
	}
}
