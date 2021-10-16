package waitroom

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
)

var ticker = time.NewTicker(time.Second)

func Exec() {
	collection.OnceCollection[ui.WAIT_ROOM].Do(func() {
		connection.UseConnection().Call("reg_user", pc.UsePC(), nil)

		connection.UseConnection().Call("reg_world", struct {
			World world.World
			PC    pc.PC
		}{
			*world.UseWorld(), *pc.UsePC(),
		}, nil)
	})

	select {
	case <-ticker.C:
		connection.UseConnection().Call("update_world_users", world.UseWorld(), &world.UseWorld().Users)
	default:
	}
	// if !l.currState.NetworkingStates.LobbyWaitRoom{
	// 	l.currState.NetworkingStates.LobbyWaitRoom = true
	// 	go func(){
	// 		parser := Server.GameParser(new(Server.GameRequest))
	// 		server := Server.Network(new(Server.N))
	// 		server.Init(nil, l.userConfig, 1, nil, parser.Parse, "GetUsersInfoPrepLobby")
	// 		server.Write()
	// 		response := server.ReadGame(parser.Unparse)
	// 		responseUser :=  GetUserFromList(l.userConfig.PersonalInfo.Username, response)

	// 		if responseUser != nil{
	// 			switch responseUser.Error{
	// 			case "60":
	// 				iswritten := func(u string)bool{
	// 					for _, us := range l.winConf.WaitRoom.NewMembers{
	// 						if u == us{
	// 							return true
	// 						}
	// 					}
	// 					return false
	// 				}
	// 				for _, user := range response{
	// 					if len(l.winConf.WaitRoom.NewMembers) <= 4 && !iswritten(user.PersonalInfo.Username){
	// 						l.winConf.WaitRoom.NewMembers = append(l.winConf.WaitRoom.NewMembers, user.PersonalInfo.Username)
	// 					}
	// 				}
	// 			case "502":
	// 				l.currState.MainStates.SetGame()
	// 			}
	// 		}
	// 		l.currState.NetworkingStates.LobbyWaitRoom = false
	// 	}()
	// }
}
