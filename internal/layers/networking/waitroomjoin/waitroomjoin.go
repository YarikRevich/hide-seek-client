package waitroomjoin

import (
	"context"
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/latency"
	"github.com/YarikRevich/HideSeek-Client/internal/core/middlewares"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_JOIN, func() {
		pm := objects.UseObjects().PC().ToAPIMessage()
		networking.UseNetworking().Dialer().Conn().AddPC(context.Background(), pm, grpc.EmptyCallOption{})
		
		w := objects.UseObjects().World()
		r, err := networking.UseNetworking().Dialer().Conn().GetWorld(context.Background(), &api.GetWorldRequest{WorldId: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil{
			logrus.Fatal(err)
		}
		w.FromAPIMessage(r)
	})
	
	latency.UseLatency().Timings().ExecEach(func() {
		w := objects.UseObjects().World()
		worldId := w.ID.String()
	
		worldObjects, err := networking.UseNetworking().Dialer().Conn().GetWorldObjects(context.Background(), &api.WorldObjectsRequest{WorldId: worldId}, grpc.EmptyCallOption{})
		if err != nil{
			logrus.Fatal(err)
		}

		w.UpdateObjects(worldObjects)

		r, err := networking.UseNetworking().Dialer().Conn().IsGameStarted(context.Background(), &api.IsGameStartedRequest{WorldId: worldId}, grpc.EmptyCallOption{}); 
		if err != nil{
			logrus.Fatal(err)
		}
		
		if r.Started{
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})
			statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_GAME)
		}
	}, statemachine.UI_WAIT_ROOM_JOIN, time.Millisecond * 300)
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
	// latency.UseLatency().Timings().ExecEach(func(){
	// 	w := objects.UseObjects().World()
	// 	networking.UseNetworking().Dialer().Conn().Call("update_world", w.ID, w)
	// }, time.Second)

}
