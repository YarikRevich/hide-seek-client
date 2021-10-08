package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/render/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/syncer"
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
	syncer.SyncConfValues(screen)

	for _, dt := range render.GetToRender(){
		dt(screen)
	}
	middlewares.UseRenderMiddlewares()
	render.CleanRenderPool()

	w, h := screen.Size()
	history.SetScreenSize(history.ScreenSize{
		Height: h, Width: w})
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
