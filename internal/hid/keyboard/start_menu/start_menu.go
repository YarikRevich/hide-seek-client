package startmenu

import (
	"fmt"
	"github.com/YarikRevich/HideSeek-Client/internal/buffers/text"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Exec(){
	for _, v := range inpututil.PressedKeys() {
		text.UseBuffer().Write(v.String())
	}

	render.SetTextToRender(func(screen *ebiten.Image) {
			ebitenutil.DebugPrint(screen,
				fmt.Sprintf(
					"%s\n",
					text.UseBuffer().Read()))
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