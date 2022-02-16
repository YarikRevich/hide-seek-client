package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
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
	// sl.UIManager.AddComponent(ui.NewBackground(&ui.BackgroundOpts{
	// 	ID:      "joingame",
	// 	Tilemap: sources.GetTileMap("system/backgrounds/background"),
	// }))
	// sl.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	OnMousePress: func() {
	// 		resp, err := sl.opts.NetworkingManager.ServerClient.CreateSession(context.Background(), &emptypb.Empty{})
	// 		if err != nil {
	// 			sl.opts.NotificationManager.Write("Session was not created!", -1, notifications.Error)
	// 			return
	// 		}
	// 		sl.opts.WorldManager.ID = uuid.MustParse(resp.ID)
	// 		statemachine.Layers.SetState(statemachine.LAYERS_MAP_CHOOSE)
	// 	},
	// }))
	// screenAxis := sl.opts.ScreenManager.GetAxis()

	// // fmt.Println(screenAxis.X, sl.opts.ScreenManager.GetSize())
	// sl.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	TextOpts: ui.TextOpts{
	// 		Align: sources.Center,
	// 		Text:  "gjerkgejglerjglkeglkglegeglkejgeglkgl",
	// 		Font:  sources.GetFont("base", 20),
	// 	},

	// 	ID:              "startgamebutton",
	// 	Tilemap:         sources.GetTileMap("system/buttons/button"),
	// 	SurfacePosition: types.Vec2{X: screenAxis.X, Y: screenAxis.Y - 150},
	// 	Scale:           types.Vec2{X: 4, Y: 4},

	// 	OnMousePress: func() {
	// 		statemachine.Layers.SetState(statemachine.LAYERS_SESSION)
	// 	},
	// }))
}

func NewSessionLayer() Layer {
	return &SessionLayer{
		UIManager: ui.NewUIManager(),
	}
}
