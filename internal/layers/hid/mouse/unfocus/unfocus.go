package unfocus

import "github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

func Exec() {
	statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
}
