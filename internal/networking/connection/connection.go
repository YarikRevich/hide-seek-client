package connection

import (
	"os"
	"os/signal"

	"github.com/YarikRevich/game-networking/pkg/client"
	gamenetworkingconfig "github.com/YarikRevich/game-networking/pkg/config"
	"github.com/sirupsen/logrus"
)

var instance client.Dialer

func UseConnection()client.Dialer{
	if instance == nil{
		d := client.Dial(gamenetworkingconfig.Config{
			IP: "127.0.0.1",
			Port: "8090",
		})
		instance = d

		go func(){
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Interrupt)
			for range sc{
				if err := instance.Close(); err != nil{
					logrus.Fatal(err)
				}
			}	
		}()
	}
	return instance
}