package game

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	audiocollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader/collection"
)

func Exec() {
	fmt.Println()
	switch world.UseWorld().Location.Name{
	case "helloween":
		audiocollection.GetAudioController("assets/audio/game").Start()
	case "starwars":
		audiocollection.GetAudioController("assets/audio/starwarsmaptheme").Start()
	}
	
}
