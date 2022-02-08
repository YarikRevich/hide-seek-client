package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

var Layers = []Layer{
	NewStartMenuLayer(),
}

type Layer interface {
	SetContext(*world.WorldManager)

	OnMouse()

	OnKeyboard()

	IsActive() bool
	Update()
	Render(*screen.ScreenManager)
}
