package audio

import (
	"github.com/YarikRevich/HideSeek-Client/internal/history/audio"
	audiocollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"
)

func stopPreviousTrack() {
	if p := audio.GetLastAudioTrackPath(); p != "" {
		audiocollection.GetAudioController(p).Stop()
	}
}

func UseAudioMiddleware(c func()) {
	stopPreviousTrack()
	c()
}
