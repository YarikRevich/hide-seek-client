package map_choose

import (
	// mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	// mousewheel "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_wheel"
	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func Exec() bool {
	m := events.UseEvents().Mouse()
	m.UpdateMouseWheelOffsets()
	if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().UI().SetState(ui.START_MENU),
			uimiddleware.UseUIMiddleware,
		)
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Input().SetState(input.EMPTY),
			inputmiddleware.UseInputMiddleware,
		)
		return true
	}

	for k, v := range map[string]string{
		"assets/images/maps/thumbnails/helloween": "assets/images/maps/helloween/background/background",
		"assets/images/maps/thumbnails/starwars" : "assets/images/maps/starwars/background/background",
	} {
		if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata(k)) {
			objects.UseObjects().World().SetSkin(v)
			p := objects.UseObjects().PC()
			p.InitUsername()
			p.SetSkin("assets/images/heroes/pumpkin")
			// pc.UsePC().SetSpawn(world.UseWorld().Metadata.Spawns)

			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().UI().SetState(ui.WAIT_ROOM),
				uimiddleware.UseUIMiddleware,
			)
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Input().SetState(input.EMPTY),
				inputmiddleware.UseInputMiddleware,
			)
			return true
		}
	}

	// if m.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/maps/thumbnails/starwars")) {
	// 	objects.UseObjects().World().SetSkin("assets/images/maps/starwars/background/background")
	// 	p := objects.UseObjects().PC()
	// 	p.InitUsername()
	// 	p.SetSkin("assets/images/heroes/pumpkin")
	// 	// pc.UsePC().SetSpawn(world.UseWorld().Metadata.Spawns)

	// 	applyer.ApplyMiddlewares(
	// 		statemachine.UseStateMachine().UI().SetState(ui.WAIT_ROOM),
	// 		uimiddleware.UseUIMiddleware,
	// 	)
	// 	applyer.ApplyMiddlewares(
	// 		statemachine.UseStateMachine().Input().SetState(input.EMPTY),
	// 		inputmiddleware.UseInputMiddleware,
	// 	)
	// 	return true
	// }
	return false
}
