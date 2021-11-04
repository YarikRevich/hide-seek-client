package middlewares

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/networking/connection"
	isconnect "github.com/alimasyhur/is-connect"
)

type Prepare struct{}

func (p *Prepare) Use() {
	if isconnect.IsOnline() && connection.UseConnection().IsConnected() {
		statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_ONLINE)
	}
}

func NewPrepare() *Prepare {
	return new(Prepare)
}
