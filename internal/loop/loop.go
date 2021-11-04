package loop

import (
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/syncer"
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	render.UseRender().CleanRenderPool()

	ui.Process()
	mouse.Process()
	keyboard.Process()
	networking.Process()

	if !cli.IsWithoutSound() {
		audio.Process()
	}

	return nil
}

func (g *Loop) Draw(screen *ebiten.Image) {
	screenhistory.SetScreen(screen)
	syncer.NewSyncer().Sync()

	defer func() {
		screenhistory.SetLastScreenSize()
		render.UseRender().PostRender()
	}()
	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.RENDER)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
		}()
	}
	render.UseRender().Render()
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
