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
		c.Wrap("assets/audio/game")
		c.Start()
	case "starwars":
		c.Wrap("assets/audio/starwarsmaptheme")
		c.Start()
	}
}
