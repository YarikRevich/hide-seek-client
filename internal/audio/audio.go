package audio

import (
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	// "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/audio_loader"
	// "github.com/sirupsen/logrus"
)



func Process() {
	switch statemachine.GetInstance().GetState() {
	case statemachine.GAME:
		
		// play, ok := audioloader.Audio["/audio/game.mp3"]
		// if !ok{
		// 	logrus.Fatal("audio is not found")
		// }
		// play()
	}
}
