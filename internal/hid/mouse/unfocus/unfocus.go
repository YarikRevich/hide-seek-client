package unfocus

import (
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
)

func Exec() {
	applyer.ApplyMiddlewares(
		statemachine.UseStateMachine().Input().SetState(input.EMPTY),
		inputmiddleware.UseInputMiddleware,
	)

}
