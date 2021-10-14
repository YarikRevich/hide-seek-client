package waitroom

import (
	mousepress "github.com/YarikRevich/HideSeek-Client/internal/detectors/mouse_press"
	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/input"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	inputmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/input"
	uimiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/ui"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/world"
)

func Exec()bool{
	if mousepress.IsMousePressLeftOnce(*metadatacollection.GetMetadata("assets/images/system/buttons/back")) {
		world.UseWorld().Reset()
		
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().Input().SetState(input.EMPTY),
			inputmiddleware.UseInputMiddleware,
		)
		applyer.ApplyMiddlewares(
			statemachine.UseStateMachine().UI().SetState(ui.START_MENU),
			uimiddleware.UseUIMiddleware,
		)
		return true
	}
	return false
	// if l.winConf.WaitRoom.RoomType == "create" && (l.winConf.Win.MousePosition().X >= 361 && l.winConf.Win.MousePosition().X <= 596) && (l.winConf.Win.MousePosition().Y >= 73 && l.winConf.Win.MousePosition().Y <= 165){
	// 	l.winConf.DrawWaitRoomPressedButton()
	// }

	// server := Server.Network(new(Server.N))
	// if (l.winConf.Win.MousePosition().X >= 361 && l.winConf.Win.MousePosition().X <= 596) && (l.winConf.Win.MousePosition().Y >= 73 && l.winConf.Win.MousePosition().Y <= 165) && l.winConf.Win.JustPressed(pixelgl.MouseButtonLeft){
	// 	if l.winConf.WaitRoom.RoomType == "create"{
	// 		parser := Server.GameParser(new(Server.GameRequest))
	// 		server.Init(nil, l.userConfig, 1, nil, parser.Parse, "ClosePreparingLobby")
	// 		server.Write()
	// 		server.ReadGame(parser.Unparse)
	// 		l.currState.MainStates.SetGame()
	// 	}
	// }
	// if (l.winConf.Win.MousePosition().X >= 21 && l.winConf.Win.MousePosition().X <= 68) && (l.winConf.Win.MousePosition().Y >= 463 && l.winConf.Win.MousePosition().Y <= 507) && l.winConf.Win.JustPressed(pixelgl.MouseButtonLeft){
	// 	if l.winConf.WaitRoom.RoomType == "create"{
	// 		l.winConf.WaitRoom.NewMembers = []string{}
	// 		l.currState.MainStates.SetCreateLobbyMenu()
				 
	// 		parser := Server.GameParser(new(Server.GameRequest))
	// 		server.Init(nil, l.userConfig, 1, nil, parser.Parse, "DeleteLobby")
	// 		server.Write()
	//  		server.ReadGame(parser.Unparse)
	// 	}else{
	// 		l.winConf.WaitRoom.NewMembers = []string{}
	//  		l.currState.MainStates.SetJoinLobbyMenu()
	// 	}
	// }
	// l.winConf.WindowUpdation.WaitRoomFrame++
}