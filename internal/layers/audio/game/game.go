package game

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/audiocontroller"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

func Exec() {
	c := audiocontroller.UseAudioController()
	worldMap := world.UseWorld().GetWorldMap()
	switch worldMap.Name {
	case "helloween":
		c.Wrap("game")
		c.Start()
	case "starwars":
		c.Wrap("starwarsmaptheme")
		c.Start()
	}
}
