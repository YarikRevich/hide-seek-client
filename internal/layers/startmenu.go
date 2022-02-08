package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/ui"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

type StartMenuLayer struct {
	*world.WorldManager
}

func (sml *StartMenuLayer) SetContext(w *world.WorldManager) {
	sml.WorldManager = w
}

func (sml *StartMenuLayer) IsActive() bool {
	return statemachine.Layers.Check(statemachine.LAYERS_START_MENU)
}

func (sml *StartMenuLayer) OnMouse()

func (sml *StartMenuLayer) OnKeyboard()

func (sml *StartMenuLayer) Update() {}

func (sml *StartMenuLayer) Render(sm *screen.ScreenManager) {
	ui.NewButton(&ui.ButtonOpts{
		OnMousePress: func() {
			networking.UseNetworking().Clients().Base().GetClient().CreateSession()
		},
	}).Render(sm)
}

func NewStartMenuLayer() Layer {
	return new(StartMenuLayer)
}
