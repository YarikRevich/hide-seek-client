package particles

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/layers/particles/game"
)

func Process() {
	switch statemachine.UseStateMachine().UI().GetState() {
	case statemachine.UI_GAME:
		game.Draw()
	}
}
