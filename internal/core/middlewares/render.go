package middlewares



import (
	"math"
	"sync"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking/connection"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	networkingmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
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

		if !isconnect.IsOnline() || !connection.UseConnection().IsConnected() {
			notifications.PopUp.WriteError("Servers are offline!")
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.OFFLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
		} else {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.ONLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
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