package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/YarikRevich/hide-seek-client/internal/core/ui"
)

type SessionLayer struct {
	opts      *ContextOpts
	UIManager *ui.UIManager
}

func (sl *SessionLayer) SetContext(opts *ContextOpts) {
	sl.opts = opts
}

func (sl *SessionLayer) IsActive() bool {
	return statemachine.Layers.Check(statemachine.LAYERS_SESSION)
}

func (sl *SessionLayer) Update() {
	sl.UIManager.Update(sl.opts.ScreenManager)
	sl.opts.WorldManager.Update()
}

func (sl *SessionLayer) Render() {
	sl.UIManager.Render(sl.opts.ScreenManager)
	sl.opts.WorldManager.Render(sl.opts.ScreenManager)
}

func (sl *SessionLayer) Clear() {
	sl.UIManager.Clear()
}

func (sl *SessionLayer) Init() {
	sl.UIManager.AddComponent(ui.NewBackground(&ui.BackgroundOpts{
		ID:                     "joingame",
		Tilemap:                "test/test",
		Scale:                  types.Vec2{X: 1, Y: 1},
		OrthigraphicProjection: true,

		CameraAngle:    sl.opts.WorldManager.Camera.Angle,
		CameraPitch:    sl.opts.WorldManager.Camera.Pitch,
		CameraPosition: sl.opts.WorldManager.Camera.Position,
	}))
}

func NewSessionLayer() Layer {
	return &SessionLayer{
		UIManager: ui.NewUIManager(),
	}
}
