package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	return nil
}

func (g *Loop) Draw(screen *ebiten.Image) {
	ui.Process(screen)
	mouse.Process()
	keyboard.Process()
	audio.Process()
	networking.Process()
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
