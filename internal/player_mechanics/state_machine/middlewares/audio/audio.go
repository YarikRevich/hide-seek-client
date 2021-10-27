package audio

import audiocollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"

func stopPreviousTrack() {
	if p := audiocollection.GetLastAudioTrackPath(); p != "" {
		audiocollection.GetAudioController(p).Stop()
	}
}

func UseAudioMiddleware(c func()) {
	stopPreviousTrack()
	c()
}
