package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/ui"
)

type StartMenuLayer struct {
	opts      *ContextOpts
	UIManager *ui.UIManager
}

func (sml *StartMenuLayer) SetContext(opts *ContextOpts) {
	sml.opts = opts
}

func (sml *StartMenuLayer) IsActive() bool {
	return statemachine.Layers.Check(statemachine.LAYERS_START_MENU)
}

func (sml *StartMenuLayer) Update() {
	// sml.UIManager.Update()
}

func (sml *StartMenuLayer) Render() {
	sml.UIManager.Render(sml.opts.ScreenManager)
	sml.UIManager.Clear()
}

func (sml *StartMenuLayer) Init() {
	sml.UIManager.AddComponent(ui.NewBackground(&ui.BackgroundOpts{
		Tilemap: sources.GetTileMap("system/backgrounds/background"),
	}))
	// sml.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	OnMousePress: func() {
	// 		resp, err := sml.opts.NetworkingManager.ServerClient.CreateSession(context.Background(), &emptypb.Empty{})
	// 		if err != nil {
	// 			sml.opts.NotificationManager.Write("Session was not created!", -1, notifications.Error)
	// 			return
	// 		}
	// 		sml.opts.WorldManager.ID = uuid.MustParse(resp.ID)
	// 		statemachine.Layers.SetState(statemachine.LAYERS_MAP_CHOOSE)
	// 	},
	// }))

	// sml.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	OnMousePress: func() {
	// 		statemachine.Layers.SetState(statemachine.LAYERS_MAP_CHOOSE)
	// 	},
	// }))
}

func NewStartMenuLayer() Layer {
	return &StartMenuLayer{
		UIManager: ui.NewUIManager(),
	}
}
