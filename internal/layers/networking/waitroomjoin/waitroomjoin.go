package waitroomjoin

import (
	"context"
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Exec() {
	latency.UseLatency().Once().ExecOnce(statemachine.UI_WAIT_ROOM_JOIN, func() {
		conn := networking.UseNetworking().Dialer().Conn()

		pcMess := world.UseWorld().GetPC().ToAPIMessage()

		r, err := conn.AddPC(context.Background(), pcMess, grpc.EmptyCallOption{})
		if !r.GetOk() || err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}
	})

	latency.UseLatency().Timings().ExecEach(func() {
		w := world.UseWorld()
		worldId := w.ID.String()

		conn := networking.UseNetworking().Dialer().Conn()
		worldObjects, err := conn.GetWorld(context.Background(), &wrappers.StringValue{Value: worldId}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(worldObjects.World)

		w.Update(worldObjects)

		if !w.GetGameSettings().IsWorldExist {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
			})
			notifications.PopUp.WriteError("The room was closed by the creator")
		}
		if w.GetGameSettings().IsGameStarted {
			middlewares.UseMiddlewares().UI().UseAfter(func() {
				statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
			})
		}
	}, statemachine.UI_WAIT_ROOM_JOIN, time.Second)
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
