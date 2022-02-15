package layers

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/ui"
)

type JoinMenuLayer struct {
	opts      *ContextOpts
	UIManager *ui.UIManager
}

func (iml *JoinMenuLayer) Clear() {
	iml.UIManager.Clear()
}

func (iml *JoinMenuLayer) SetContext(opts *ContextOpts) {
	iml.opts = opts
}

func (iml *JoinMenuLayer) IsActive() bool {
	return statemachine.Layers.Check(statemachine.LAYERS_START_MENU)
}

func (iml *JoinMenuLayer) Update() {
	iml.UIManager.Update()
}

func (iml *JoinMenuLayer) Render() {
	iml.UIManager.Render(iml.opts.ScreenManager)
}

func (iml *JoinMenuLayer) Init() {
	// iml.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	OnMousePress: func() {
	// 		ID, err := iml.opts.NetworkingManager.ServerClient.CreateSession()
	// 		if err != nil {
	// 			iml.opts.NotificationManager.Write("Session was not created!", -1, notifications.Error)
	// 			return
	// 		}
	// 		iml.opts.WorldManager.ID = ID
	// 		statemachine.Layers.SetState(statemachine.LAYERS_MAP_CHOOSE)
	// 	},
	// }))

	// iml.UIManager.AddComponent(ui.NewButton(&ui.ButtonOpts{
	// 	OnMousePress: func() {
	// 		statemachine.Layers.SetState(statemachine.LAYERS_MAP_CHOOSE)
	// 	},
	// }))

	// _, err := sml.opts.NetworkingManager.ServerClient.FindSession()
	// if err != nil {
	// 	sml.opts.NotificationManager.Write("Session was not found!", -1, notifications.Error)
	// 	return
	// }
	// sml.opts.WorldManager.ID =
	// return iml
}

func NewJoinMenuLayer() Layer {
	return &JoinMenuLayer{
		UIManager: ui.NewUIManager(),
	}
}

// _, err := iml.opts.NetworkingManager.ServerClient.FindSession()
// if err != nil {
// 	iml.opts.NotificationManager.Write("Session was not found!", -1, notifications.Error)
// 	return
// }
// iml.opts.WorldManager.ID =
