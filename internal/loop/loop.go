package loop

import (
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/syncer"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/ui"
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
	if !cli.GetDebug() {
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
	if cli.GetDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.RENDER)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
			// profiling.UseProfiler().SumUpMonitoring()
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
