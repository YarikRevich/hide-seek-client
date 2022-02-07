package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
)

type StartMenuLayer struct {
}

func (sml *StartMenuLayer) SetContext() {

}

func (sml *StartMenuLayer) IsActive() bool {
	return statemachine.Layers.Check(statemachine.LAYERS_START_MENU)
}

func (sml *StartMenuLayer) OnMouse() {

}

func (sml *StartMenuLayer) OnKeyboard() {

}

func (sml *StartMenuLayer) Update() {
	select {
	case <-r.ticker.C:
		r.checkServersConnectivity()
	default:
	}
}
func (sml *StartMenuLayer) Render(sm *screen.ScreenManager) {
}

func NewStartMenuLayer() Layer {
	return new(StartMenuLayer)
}
