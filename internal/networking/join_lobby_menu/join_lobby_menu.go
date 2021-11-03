package joinlobbymenu

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
)

var ticker = time.NewTicker(time.Second)

func Exec() {

	// if j.currState.SendStates.JoinRoom {

	// 	j.userConfig.PersonalInfo.LobbyID = strings.Join(j.winConf.TextAreas.JoinLobbyInput.WrittenText, "")

	// 	parser := Server.GameParser(new(Server.GameRequest))
	// 	server := Server.Network(new(Server.N))
	// 	server.Init(nil, j.userConfig, 1, nil, parser.Parse, "AddToLobby")
	// 	server.Write()
	// 	response := server.ReadGame(parser.Unparse)
	// 	switch response[0].Error {
	// 	case "20":
	// 		j.winConf.WaitRoom.RoomType = "join"
	// 		j.currState.MainStates.SetWaitRoom()

	// 	case "500":
	// 		j.winConf.WindowError.LobbyDoesNotExist = true
	// 		j.winConf.WindowError.LobbyErrorStop = time.Now()
	// 		j.winConf.WindowError.LobbyErrorText = "Such lobby doesn't exist!"
	// 	}
	// 	j.currState.SendStates.JoinRoom = false
	// }
	select {
	case <-ticker.C:
		w := objects.UseObjects().World()
		connection.UseConnection().Call("update_world", w.ID, w)
	default:
	}
}
