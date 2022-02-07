package middlewares

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	isconnect "github.com/alimasyhur/is-connect"
)

// type Render struct {
// 	sync.Mutex

// 	ticker *time.Ticker
// }

// func (r *Render) cleanPopUp() {

// }

func (r *Render) checkServersConnectivity() {
	go func() {
		r.Lock()

		if !isconnect.IsOnline() && statemachine.UseStateMachine().Dial().GetState() == statemachine.DIAL_WAN {
			notifications.PopUp.WriteError("You are offline, turn on LAN server to play locally!")
			us := statemachine.UseStateMachine().UI().GetState()
			if !(us == statemachine.UI_SETTINGS_MENU || us == statemachine.UI_START_MENU) {
				statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_OFFLINE)
			}
		} else {
			if !networking.UseNetworking().Dialer().IsServerClientConnected() {
				networking.UseNetworking().Dialer().ReconnectServerClient()

				notifications.PopUp.WriteError("Servers are offline!")
				statemachine.Networking.SetState(statemachine.NETWORKING_OFFLINE)

			} else {
				statemachine.Networking.SetState(statemachine.NETWORKING_ONLINE)
			}
		}

		r.Unlock()
	}()
}

// func (r *Render) UseAfter(c func()) {
// // 	c()

// // }

// func NewRender() *Render {
// 	return &Render{ticker: time.NewTicker(3 * time.Second)}
// }
