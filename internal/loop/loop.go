package loop

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/player"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/ingame"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/YarikRevich/hide-seek-client/internal/layers"

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
	*screen.ScreenManager
	*world.WorldManager
	*notifications.NotificationManager
}

var _ ebiten.Game = (*Loop)(nil)

func (g *Loop) Update() error {
	if !params.IsWithoutSound() {
		audio.Process()
	}

	mouse.Process()
	screen.UseScreen().CleanScreen()

	networking.Process()

	ui.Process()
	particles.Process()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Update()
	}

	if params.IsWithoutSound() {
		player.UsePlayer().StopAll()
	}

	// transition.UseTransitionPool().Process()

	keyboard.Process()

	return nil
}

func (g *Loop) Draw(i *ebiten.Image) {
	g.ScreenManager.SetImage(i)
	// screen.UseScreen().SetScreen(i)

	if params.IsDebug() {
		ingame.UseProfiler().StartMonitoring(ingame.RENDER)
		defer ingame.UseProfiler().StopMonitoring(ingame.RENDER)
	}

	for _, v := range layers.Layers {
		if v.IsActive() {
			v.SetContext(g.WorldManager)
			v.Update()
			v.Render(g.ScreenManager)
		}
	}

	// screen.UseScreen().SetLastSize()

	if params.IsDebug() {
		debugui.UseDebugImGUI().Render(i)
	}
}

func (g *Loop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New(sm *screen.ScreenManager) *Loop {
	return &Loop{
		ScreenManager:       sm,
		WorldManager:        world.NewWorldManager(),
		NotificationManager: notifications.NewNotificationManager()}
}
