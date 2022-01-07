package game

import (
	"context"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// "time"

// "github.com/YarikRevich/hide-seek-client/internal/core/latency"
// "github.com/YarikRevich/hide-seek-client/internal/core/networking"
// "github.com/YarikRevich/hide-seek-client/internal/core/objects"

// "github.com/YarikRevich/hide-seek-client/internal/gameplay/world"
// "github.com/YarikRevich/hide-seek-client/internal/networking/collection"
// "github.com/YarikRevich/hide-seek-client/internal/networking/connection"
// "github.com/YarikRevich/hide-seek-client/internal/player_mechanics/state_machine/constants/ui"
// "github.com/hajimehoshi/ebiten/v2"
// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
// "github.com/sirupsen/logrus"

func Exec() {
	latency.UseLatency().Timings().ExecEach(func() {
		// w := world.UseWorld()
		// for _, pc := range w.GetPCs() {
		// 	if statemachine.UseStateMachine().PCs().GetState(pc.ID) == statemachine.PC_DEAD_NOW {
		// 		// base := networking.UseNetworking().Clients().Base().GetClient()
		// 	}
		// }
		w := world.UseWorld()
		p := w.GetPC()
		client := networking.UseNetworking().Clients().Base().GetClient()

		if !w.GetGameSettings().IsWorldExist {
			if _, err := client.DeleteWorld(context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{}); err != nil {
				notifications.PopUp.WriteError(err.Error())
				return
			}
		}

		if _, err := client.InsertOrUpdateWorld(context.Background(), w.ToAPIMessage(), grpc.EmptyCallOption{}); err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}

		if _, err := client.InsertOrUpdatePC(context.Background(), p.ToAPIMessage()); err != nil {
			notifications.PopUp.WriteError(err.Error())
			return
		}

		worldObjects, err := client.FindWorldObjects(
			context.Background(), &wrappers.StringValue{Value: w.ID.String()}, grpc.EmptyCallOption{})
		if err != nil {
			logrus.Fatal(err)
		}

		w.Update(worldObjects)

		// if p.IsKicked {
		// 	if _, err := client.DeletePC(context.Background(), &wrappers.StringValue{Value: p.Base.ID.String()}); err != nil {
		// 		logrus.Fatal(err)
		// 	}
		// 	middlewares.UseMiddlewares().UI().UseAfter(func() {
		// 		statemachine.UseStateMachine().UI().SetState(statemachine.UI_START_MENU)
		// 	})
		// 	notifications.PopUp.WriteError("You were kicked from the session")
		// 	p.SetKicked(false)
		// }
	}, statemachine.UI_GAME, time.Second)
	// 	w := objects.UseObjects().World()
	// 	p := objects.UseObjects().PC()
	// 	networking.UseNetworking().Dialer().client().Call("update_game", w, p.ID)
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
