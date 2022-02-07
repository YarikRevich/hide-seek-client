package layers

import "github.com/YarikRevich/hide-seek-client/internal/core/screen"

var Layers = []Layer{
	NewStartMenuLayer(),
}

type Layer interface {
	SetContext(world.WorldManager)

	IsActive() bool
	Update()
	Render(*screen.ScreenManager)
}
