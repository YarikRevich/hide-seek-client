package game

import (
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/keyboard"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/mouse"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/networking"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/screen"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/sound"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var _ ebiten.Game = (*Game)(nil)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.Process()
	mouse.Process()
	keyboard.Process()
	sound.Process()
	networking.Process()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	switch screen.GetInstance().GetState() {
	case screen.FULLSCREEN:
		return ebiten.ScreenSizeInFullscreen()
	}
	return outsideWidth, outsideHeight
}

func New() *Game {
	return new(Game)
}
