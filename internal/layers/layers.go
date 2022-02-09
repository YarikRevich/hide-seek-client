package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
)

var Layers = []Layer{
	NewStartMenuLayer().Init(),
}

type ContextOpts struct {
	ScreenManager       *screen.ScreenManager
	WorldManager        *world.WorldManager
	NotificationManager *notifications.NotificationManager
	NetworkingManager   *networking.NetworkingManager
}

type Layer interface {
	SetContext(*ContextOpts)

	Init() Layer

	IsActive() bool
	Update()
	Render()
}
