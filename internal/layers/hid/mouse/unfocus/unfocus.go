package unfocus

import "github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

func Exec() {
	statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
}
