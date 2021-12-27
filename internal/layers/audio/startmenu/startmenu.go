package startmenu

import "github.com/YarikRevich/hide-seek-client/internal/core/audiocontroller"

func Exec() {
	c := audiocontroller.UseAudioController()
	c.Wrap("background")
	c.Start()
}
