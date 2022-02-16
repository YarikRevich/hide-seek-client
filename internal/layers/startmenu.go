package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
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
	sml.UIManager.Update(sml.opts.ScreenManager)
	sml.opts.WorldManager.Update()
}

func (sml *StartMenuLayer) Render() {
	sml.UIManager.Render(sml.opts.ScreenManager)
	sml.opts.WorldManager.Render(sml.opts.ScreenManager)
}

func (sml *StartMenuLayer) Clear() {
	sml.UIManager.Clear()
}

func (sml *StartMenuLayer) Init() {
	sml.UIManager.AddComponent(ui.NewBackground(&ui.BackgroundOpts{
		ID:      "joingame",
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
	screenAxis := sml.opts.ScreenManager.GetAxis()

	// fmt.Println(screenAxis.X, sml.opts.ScreenManager.GetSize())
	sml.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
		TextOpts: ui.TextOpts{
			Align: sources.Center,
			Text:  "gjerkgejglerjglkeglkglegeglkejgeglkgl",
			Font:  sources.GetFont("base", 20),
		},

		ID:              "startgamebutton",
		Tilemap:         sources.GetTileMap("system/buttons/button"),
		SurfacePosition: types.Vec2{X: screenAxis.X, Y: screenAxis.Y - 150},
		Scale:           types.Vec2{X: 4, Y: 4},

		OnMousePress: func() {
			statemachine.Layers.SetState(statemachine.LAYERS_SESSION)
		},
	}))
}

func NewStartMenuLayer() Layer {
	return &StartMenuLayer{
		UIManager: ui.NewUIManager(),
	}
}
