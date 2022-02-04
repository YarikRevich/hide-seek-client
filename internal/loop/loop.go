package loop

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/transition"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"

	"github.com/YarikRevich/hide-seek-client/internal/layers/animation"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/keyboard"
	"github.com/YarikRevich/hide-seek-client/internal/layers/hid/mouse"
	"github.com/YarikRevich/hide-seek-client/internal/layers/networking"
	"github.com/YarikRevich/hide-seek-client/internal/layers/particles"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui"

	"github.com/YarikRevich/hide-seek-client/tools/debugui"
	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/hajimehoshi/ebiten/v2"
)

type Loop struct {
	screen.ScreenManager
	UIManager
	WorldManager
}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	render.UseRender().CleanRenderPool()

	if !params.IsWithoutSound() {
		audio.Process()
	}

	mouse.Process()
	screen.UseScreen().CleanScreen()

	networking.Process()
	animation.Process()

	ui.Process()
	particles.Process()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Update()
	}

	transition.UseTransitionPool().Process()

	keyboard.Process()

	return nil
}

func (g *Loop) Draw(i *ebiten.Image) {
	screen.UseScreen().SetScreen(i)

	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.RENDER)
		defer ingame.UseProfiler().StopMonitoring(ingame.RENDER)
	}

	middlewares.UseMiddlewares().Render().UseAfter(render.UseRender().Render)

	screen.UseScreen().SetLastSize()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Render(i)
	}
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New() *Loop {
	return &Loop{ScreenManager: screen.ScreenManager{Pixels: make([]types.Vec2, 0)}}
}
