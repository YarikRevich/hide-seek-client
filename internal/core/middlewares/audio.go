package middlewares

import "github.com/YarikRevich/HideSeek-Client/internal/core/audiocontroller"

type Audio struct{}

func (a *Audio) stopLastTrack() {
	c := audiocontroller.UseAudioController()

	if c.LastTrackPath != ""{
		c.Wrap(c.LastTrackPath)
		c.Stop()
	}
}

func (a *Audio) UseAfter(c func()){
	c()
	a.stopLastTrack()
}

func NewAudio() *Audio {
	return new(Audio)
}
