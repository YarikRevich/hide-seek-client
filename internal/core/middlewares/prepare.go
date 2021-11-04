package middlewares

import (
	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking/connection"
	isconnect "github.com/alimasyhur/is-connect"
)

type Prepare struct {}

func (p *Prepare) Use(){
	if isconnect.IsOnline() && connection.UseConnection().IsConnected() {
		UseMiddlewares().Networking().UseAfter(func(){
			statemachine.UseStateMachine().Networking().SetState(networking.ONLINE),
		})
	}
}

func NewPrepare() *Prepare{
	return new(Prepare)
}