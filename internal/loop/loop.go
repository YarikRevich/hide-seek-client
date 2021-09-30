package loop

import (
	// "fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"

	// "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/ai/collisions"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	ui.Process()
	mouse.Process()
	keyboard.Process()
	audio.Process()
	networking.Process()
	return nil
}

func (g *Loop) Draw(screen *ebiten.Image) {
	for _, c := range render.GetImagesToRender(){
		collisions.SyncCollisionWithImage(screen, c.Image)
		screen.DrawImage(c.Image, c.CallBack(screen))
	}
	for _, dt := range render.GetTextToRender(){
		dt(screen)
	}

	render.CleanRenderPool()
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	// switch screen.GetInstance().GetState() {
	// case screen.FULLSCREEN:
	// 	return ebiten.ScreenSizeInFullscreen()
	// }
	// maxW, maxH := ebiten.ScreenSizeInFullscreen()

	// if (int((float64(outsideWidth) / float64(maxW)) * 100) < 40 &&  int((float64(outsideHeight)/ float64(maxH)) * 100)  < 40){
		// fmt.Println("HEre", float64(outsideWidth) / float64(maxW))
		// ebiten.SetWindowResizable(false)
	// }

	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
