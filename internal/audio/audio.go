package audio

import (
	// "github.com/YarikRevich/HideSeek-Client/internal/audio/game"
	// startmenu "github.com/YarikRevich/HideSeek-Client/internal/audio/start_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/audio"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	audiomiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/audio"
)

func Process() {
	if statemachine.UseStateMachine().Audio().GetState() == audio.DONE {
		switch statemachine.UseStateMachine().UI().GetState() {
		case ui.START_MENU:
			// startmenu.Exec()
		case ui.GAME:
			// game.Exec()
		default:
			return
		}

		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Audio().SetState(audio.UNDONE),
			audiomiddleware.UseAudioMiddleware)
	}
}
