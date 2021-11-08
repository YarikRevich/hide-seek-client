package middlewares

import (
	"math"
	"sync"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	// "github.com/YarikRevich/HideSeek-Client/internal/layers/networking/connection"
	isconnect "github.com/alimasyhur/is-connect"
)

type Render struct {
	sync.Mutex

	ticker *time.Ticker
}

func (r *Render) cleanPopUp(){
	notifications.PopUp.Filter(func(e *notifications.NotificatorEntity) bool {
		return math.Signbit(float64(time.Now().Unix() - e.Timestamp))
	})
}

func (r *Render) blockRenderIfOffline(){
	go func() {
		r.Lock()

		// || !connection.UseConnection().IsConnected()
		if !isconnect.IsOnline() {
			notifications.PopUp.WriteError("Servers are offline!")
			statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_OFFLINE)
		} else {
			statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_ONLINE)
		}

		r.Unlock()
	}()
}


func (r *Render) UseAfter(c func()){
	c()

	select {
	case <- r.ticker.C:
		r.blockRenderIfOffline()
	default:
	}
	
	r.cleanPopUp()
}

func NewRender()*Render{
	return &Render{ticker: time.NewTicker(3 * time.Second)}
} 