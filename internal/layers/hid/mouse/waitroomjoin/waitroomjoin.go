package waitroomjoin


import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	
	if m.IsAnyMouseButtonsPressed() {
		if m.IsMousePressLeftOnce(*sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified) {
			objects.UseObjects().World().ResetPCs()

			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_HERO_CHOOSE)
			})

			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_EMPTY)
			return true
		}
	}

	return false
}
