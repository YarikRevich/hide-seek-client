package connection

import (
	"os"
	"os/signal"

	statemachine "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/constants/networking"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/applyer"
	networkingmiddleware "github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/middlewares/networking"
	"github.com/YarikRevich/game-networking/pkg/client"
	gamenetworkingconfig "github.com/YarikRevich/game-networking/pkg/config"
	"github.com/sirupsen/logrus"
)

var instance client.Dialer

func UseConnection() client.Dialer {
	if instance == nil {
		d, err := client.Dial(gamenetworkingconfig.Config{
			IP:   "127.0.0.1",
			Port: "8090",
		})
		if err != nil {
			applyer.ApplyMiddlewares(
				statemachine.UseStateMachine().Networking().SetState(networking.OFFLINE),
				networkingmiddleware.UseNetworkingMiddleware,
			)
		}

		instance = d

		go func() {
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Interrupt)
			for range sc {
				if err := instance.Close(); err != nil {
					logrus.Fatal(err)
				}
			}
		}()
	}
	return instance
}
