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

var ticker = time.NewTicker(time.Second * 3)
var once sync.Once
var m sync.Mutex

func isAllowedToUseMiddlewares() bool {
	select {
	case <-ticker.C:
		return true
	default:
		return false
	}
}

func checkPopUpMessagesToClean() {
	notifications.PopUp.Filter(func(e *notifications.NotificatorEntity) bool {
		return math.Signbit(float64(time.Now().Unix() - e.Timestamp))
	})
}

func checkIfOnlineInitial() {
	once.Do(func() {
		if isconnect.IsOnline() && connection.UseConnection().IsConnected() {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.ONLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
		}
	})
}

func checkIfOnline() {
	go func() {
		m.Lock()

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

		m.Unlock()
	}()
}

func UseRenderMiddlewares() {
	checkIfOnlineInitial()
	if isAllowedToUseMiddlewares() {
		checkIfOnline()
	}
	checkPopUpMessagesToClean()
}
