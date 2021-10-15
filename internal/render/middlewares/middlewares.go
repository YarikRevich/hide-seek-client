package middlewares

import (
	"math"
	"sync"
	"time"

	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	networkingmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/networking"
	popupmessagescollection "github.com/YarikRevich/HideSeek-Client/internal/pop_up_messages/collection"
	popupmessagescommon "github.com/YarikRevich/HideSeek-Client/internal/pop_up_messages/common"
	isconnect "github.com/alimasyhur/is-connect"
)

var ticker = time.NewTicker(time.Second * 1)
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
	popupmessagescollection.PopUpMessages.Filter(func(e *popupmessagescommon.PopUpEntity) bool {
		return math.Signbit(float64(time.Now().Unix() - e.Timestamp))
	})
}

func checkIfOnlineInitial() {
	once.Do(func() {
		if isconnect.IsOnline() {
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

		if !isconnect.IsOnline() {
			popupmessagescollection.PopUpMessages.WriteError("You are offline!")
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.OFFLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
			m.Unlock()
		} else {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.ONLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
		}
	}()
}

func UseRenderMiddlewares() {
	checkIfOnlineInitial()
	if isAllowedToUseMiddlewares() {
		checkIfOnline()
	}
	checkPopUpMessagesToClean()
}
