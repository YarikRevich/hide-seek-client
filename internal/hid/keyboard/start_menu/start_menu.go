package startmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/handler"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/collection"
	buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/common"

	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	handler.HandleKeyboardPress(buffercollection.SettingsMenuNameBuffer, []handler.PipelineEntity{
		{Keys: []ebiten.Key{ebiten.KeyBackspace}, Callback: func(b common.IBuffer, k rune){
			b.Pop()
		}},
		{Keys: collection.UserKeys, Callback: func(b common.IBuffer, k rune){
			b.Push(k)
		}},
	})

	// 	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 320 && s.winConf.Win.MousePosition().Y <= 415){
	// 		s.winConf.DrawStartMenuPressedCreateButton()
	// 	}

	// 	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 144 && s.winConf.Win.MousePosition().Y <= 242){
	// 		s.winConf.DrawStartMenuPressedJoinButton()
	// 	}

	// 	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 320 && s.winConf.Win.MousePosition().Y <= 415) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft) {
	// 		s.currState.MainStates.SetCreateLobbyMenu()
	// 	}
	// 	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 144 && s.winConf.Win.MousePosition().Y <= 242) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft) {
	// 		s.currState.MainStates.SetJoinLobbyMenu()
	// 	}
}
