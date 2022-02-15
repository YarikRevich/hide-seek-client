package loop

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/YarikRevich/hide-seek-client/internal/layers"

	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/layers/audio"

	"github.com/YarikRevich/hide-seek-client/internal/layers/particles"
	"github.com/YarikRevich/hide-seek-client/internal/layers/ui"

	"github.com/YarikRevich/hide-seek-client/tools/debugui"
	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/hajimehoshi/ebiten/v2"
)

type LoopOpts struct {
	ScreenManager       *screen.ScreenManager
	WorldManager        *world.WorldManager
	NotificationManager *notifications.NotificationManager
	NetworkingManager   *networking.NetworkingManager
}

type Loop struct {
	opts *LoopOpts
}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	if !params.IsWithoutSound() {
		audio.Process()
	}

	g.opts.ScreenManager.CleanScreen()

	// networking.Process()

	ui.Process()
	particles.Process()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Update()
	}

	if params.IsWithoutSound() {
		player.UsePlayer().StopAll()
	}

	for _, v := range layers.Layers {
		if v.IsActive() {
			v.Clear()
			v.Init()
			v.Update()
		}
	}

	// transition.UseTransitionPool().Process()

	return nil
}

func (g *Loop) Draw(i *ebiten.Image) {
	g.opts.ScreenManager.SetImage(i)

	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.RENDER)
		defer ingame.UseProfiler().StopMonitoring(ingame.RENDER)
	}

	for _, v := range layers.Layers {
		if v.IsActive() {
			v.Render()
		}
	}

	g.opts.ScreenManager.SetLastSize()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Render(i)
	}
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.opts.ScreenManager.SetSizeOnStartup(types.Vec2{
		X: float64(outsideWidth),
		Y: float64(outsideHeight)})
	return outsideWidth, outsideHeight
}

func New(opts *LoopOpts) *Loop {
	for _, v := range layers.Layers {
		v.SetContext(&layers.ContextOpts{
			ScreenManager:       opts.ScreenManager,
			NotificationManager: opts.NotificationManager,
			NetworkingManager:   opts.NetworkingManager,
			WorldManager:        opts.WorldManager,
		})
	}
	return &Loop{opts}
}
