package startmenu

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/buffers/text"
	"github.com/YarikRevich/HideSeek-Client/internal/cursor"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Exec() {
	b := text.UseBuffer()
	b.Pop()

	for _, v := range inpututil.PressedKeys() {
		if inpututil.KeyPressDuration(v) == 1 {
			switch v {
			case ebiten.KeyEnter:
				b.Push('\n')
			case ebiten.KeyControl:
				return
			default:
				for _, r := range v.String(){
					b.Push(r)
				}
			}
		}
	}

	cursor.SetCursorBlink(b)

	render.SetToRender(func(screen *ebiten.Image) {
		ebitenutil.DebugPrint(screen,
			fmt.Sprintf(
				"%s\n",
				b.Read()))
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
