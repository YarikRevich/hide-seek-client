package middlewares

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	isconnect "github.com/alimasyhur/is-connect"
)

type Prepare struct{}

func (p *Prepare) Use() {
	if isconnect.IsOnline() {
		statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_ONLINE)
	}
}

func NewPrepare() *Prepare {
	return new(Prepare)
}
