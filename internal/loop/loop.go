package loop

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sourceupgrader"
	"github.com/YarikRevich/HideSeek-Client/internal/core/transition"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/keyboard"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/hid/mouse"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/ui"

	"github.com/YarikRevich/HideSeek-Client/tools/debugui"
	"github.com/YarikRevich/HideSeek-Client/tools/params"
	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct{}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	screen.UseScreen().CleanScreen()
	render.UseRender().CleanRenderPool()

	mouse.Process()
	networking.Process()
	animation.Process()

	ui.Process()
	debugui.UseDebugImGUI().Update()

	transition.UseTransitionPool().Process()

	keyboard.Process()

	if !params.IsWithoutSound() {
		audio.Process()
	}

	return nil
}

func (g *Loop) Draw(i *ebiten.Image) {
	screen.UseScreen().SetScreen(i)

	sourceupgrader.NewUpgrader().Upgrade()

	if params.IsDebug() {
		profiling.UseProfiler().StartMonitoring(profiling.RENDER)
		defer profiling.UseProfiler().EndMonitoring()
	}

	middlewares.UseMiddlewares().Render().UseAfter(render.UseRender().Render)

	screen.UseScreen().SetLastSize()

	debugui.UseDebugImGUI().Render(i)
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return new(Loop)
}
