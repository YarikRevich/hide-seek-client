package loop

import (
	// "github.com/YarikRevich/HideSeek-Client/internal/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/syncer"
	"github.com/YarikRevich/HideSeek-Client/internal/ui"
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
	// audio.Process()

	return nil
}

func (g *Loop) Draw(screen *ebiten.Image) {
	profiling.UseProfiler().StartMonitoring(profiling.RENDER)

	syncer.SyncConfValues(screen)

	render.UseRender().UpdateScreen(screen)
	render.UseRender().Render()

	history.UpdateScreenSize(screen)

	profiling.UseProfiler().EndMonitoring()
	profiling.UseProfiler().SumUpMonitoring()

	render.UseRender().PostRender()
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
