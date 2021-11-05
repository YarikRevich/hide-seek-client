package startmenu

import "github.com/YarikRevich/HideSeek-Client/internal/core/audiocontroller"

func Exec() {
	c := audiocontroller.UseAudioController()
	c.Wrap("assets/audio/inter")
	c.Start()
}
