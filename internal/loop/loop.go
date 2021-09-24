package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/screen"
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
	for i, o := range render.GetImagesToRender(){
		//check collisions
		screen.DrawImage(i, o(screen))
	}
	for _, dt := range render.GetDebugTextToRender(){
		dt(screen)
	}
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	switch screen.GetInstance().GetState() {
	case screen.FULLSCREEN:
		return ebiten.ScreenSizeInFullscreen()
	}
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
