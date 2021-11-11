package game

import (
	// "time"

	// "github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	// "github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	// "github.com/YarikRevich/HideSeek-Client/internal/core/objects"
)

// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
// "github.com/YarikRevich/HideSeek-Client/internal/networking/collection"
// "github.com/YarikRevich/HideSeek-Client/internal/networking/connection"
// "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/ui"
// "github.com/hajimehoshi/ebiten/v2"
// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
// "github.com/sirupsen/logrus"

func Exec() {
	// latency.UseLatency().Timings().ExecEach(func() {
	// 	w := objects.UseObjects().World()
	// 	p := objects.UseObjects().PC()
	// 	networking.UseNetworking().Dialer().Conn().Call("update_game", w, p.ID)
	// }, time.Millisecond * 300)
	// if !g.currState.NetworkingStates.GameProcess {
	// 	g.currState.NetworkingStates.GameProcess = true
	// 	go func() {
	// 		parser := Server.GameParser(new(Server.GameRequest))
	// 		server := Server.Network(new(Server.N))
	// 		server.Init(nil, g.userConfig, 0, nil, parser.Parse, "GetUsersInfoReadyLobby")

	// 		server.Write()
	// 		response := server.ReadGame(parser.Unparse)
	// 		responseUser := GetUserFromList(g.userConfig.PersonalInfo.Username, response)
	// 		if response != nil {
	// 			switch responseUser.Error {
	// 			case "70":
	// 				cp := ConfigParsers.ConfigParser(new(ConfigParsers.CP))
	// 				cp.Init(g.winConf, g.userConfig)
	// 				cp.ApplyConfig(responseUser)
	// 				for _, value := range response {
	// 					nu := cp.Unparse(value)
	// 					cp.Commit(nu)
	// 				}
	// 			case "502":
	// 				g.currState.MainStates.SetStartMenu()
	// 			}
	// 		}
	// 		g.currState.NetworkingStates.GameProcess = false
	// 	}()
	// }
}
