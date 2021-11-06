package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/runtime"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/syncer"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui"

	"github.com/YarikRevich/HideSeek-Client/tools/cli"
	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	runtime.UseRuntime().SetPrepared()
	render.UseRender().CleanRenderPool()

	animation.Process()
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

	if cli.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.RENDER)
		defer func() {
			profiling.UseProfiler().EndMonitoring()
		}()
	}
	
	middlewares.UseMiddlewares().Render().UseAfter(render.UseRender().Render)

	screenhistory.SetLastScreenSize()
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
