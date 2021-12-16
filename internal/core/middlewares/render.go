package middlewares

import (
	"math"
	"sync"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/audiocontroller"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/tools/params"

	isconnect "github.com/alimasyhur/is-connect"
)

type Render struct {
	sync.Mutex

	ticker *time.Ticker
}

func (r *Render) cleanPopUp() {
	notifications.PopUp.Filter(func(e *notifications.NotificatorEntity) bool {
		return math.Signbit(float64(time.Now().Unix() - e.Timestamp))
	})
}

func (r *Render) blockRenderIfOffline() {
	go func() {
		r.Lock()

		if !isconnect.IsOnline() && statemachine.UseStateMachine().Dial().GetState() == statemachine.DIAL_WAN {
			notifications.PopUp.WriteError("You are offline, turn on LAN server to play locally!")
			us := statemachine.UseStateMachine().UI().GetState()
			if !(us == statemachine.UI_SETTINGS_MENU || us == statemachine.UI_START_MENU) {
				statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_OFFLINE)
			}
		} else {
			if !networking.UseNetworking().Dialer().IsConnected() {
				networking.UseNetworking().Dialer().Reconnect()

				notifications.PopUp.WriteError("Servers are offline!")
				statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_OFFLINE)

			} else {
				statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_ONLINE)
			}
		}

		r.Unlock()
	}()
}

func (r *Render) UseAfter(c func()) {
	c()

	select {
	case <-r.ticker.C:
		r.blockRenderIfOffline()
	default:
	}

	//Handles sound disabling, because in debug
	//mode sound can be disabled after music play start
	if params.IsWithoutSound() {
		audiocontroller.UseAudioController().Stop()
	}
	r.cleanPopUp()
}

func NewRender() *Render {
	return &Render{ticker: time.NewTicker(3 * time.Second)}
}
