package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	audiocollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"
)

func Exec() {
	switch objects.UseObjects().World().Name{
	case "helloween":
		audiocollection.GetAudioController("assets/audio/game").Start()
	case "starwars":
		audiocollection.GetAudioController("assets/audio/starwarsmaptheme").Start()
	}
}
