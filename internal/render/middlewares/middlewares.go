package middlewares

import (
	"math"
	"sync"
	"time"

	popupmessagescollection "github.com/YarikRevich/HideSeek-Client/internal/pop_up_messages/collection"
	popupmessagescommon "github.com/YarikRevich/HideSeek-Client/internal/pop_up_messages/common"
	isconnect "github.com/alimasyhur/is-connect"
)

var ticker = time.NewTicker(time.Second * 3)
var m sync.Mutex

func isAllowedToUseMiddlewares()bool{
	select {
	case <- ticker.C:
		return true
	default:
		return false
	}
}

func checkPopUpMessagesToClean(){
	popupmessagescollection.PopUpMessages.Filter(func(e *popupmessagescommon.PopUpEntity) bool {
		return math.Signbit(float64(time.Now().Unix() - e.Timestamp))
	})
}

func checkIfOnline(){
	go func(){
		m.Lock()

		if !isconnect.IsOnline() {
			popupmessagescollection.PopUpMessages.WriteError("TEST")
			m.Unlock()
		}
	}()
}

func UseRenderMiddlewares(){
	if isAllowedToUseMiddlewares(){
		checkIfOnline()
	}
	checkPopUpMessagesToClean()
}