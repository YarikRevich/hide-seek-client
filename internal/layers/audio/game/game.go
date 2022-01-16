package game

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

func Exec() {
	p := player.UsePlayer()
	worldMap := world.UseWorld().GetWorldMap()
	if worldMap.MetadataModel.Type.Contains("helloween") {
		p.Play("helloween", player.PlayerOpts{Infinite: true})
	} else if worldMap.MetadataModel.Type.Contains("starwars") {
		p.Play("starwars", player.PlayerOpts{Infinite: true})
	}
}
