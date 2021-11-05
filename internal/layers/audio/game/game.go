package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/audiocontroller"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
)

func Exec() {
	c := audiocontroller.UseAudioController()
	switch objects.UseObjects().World().Name{
	case "helloween":
		c.Wrap("assets/audio/game")
		c.Start()
	case "starwars":
		c.Wrap("assets/audio/starwarsmaptheme")
		c.Start()
	}
}
