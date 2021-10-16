package game

import audiocollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"

func Exec() {
	audiocollection.GetAudioController("assets/audio/game").Start()
}
