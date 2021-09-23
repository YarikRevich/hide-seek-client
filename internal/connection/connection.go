
package connection

// import (
// 	"log"

// 	"github.com/YarikRevich/game-networking/client/pkg/config"
// 	"github.com/YarikRevich/game-networking/client/pkg/connector"
// 	"github.com/YarikRevich/game-networking/client/pkg/establisher"
// 	"github.com/BurntSushi/toml"
// 	"github.com/YarikRevich/Hide-Seek-with-Guns/tools/creators"
// 	"github.com/go-ping/ping"
// )

// var (
// 	conn *establisher.Establisher
// )

// type Address struct {
// 	Addr string `toml:"addr"`
// 	Port string `toml:"port"`
// }

// type NetworkingConfig []Address

// type AddressWithRtt struct {
// 	Address
// 	AvgRtt int64
// }

// func GetServersToConnect() NetworkingConfig {
// 	var nc NetworkingConfig
// 	if _, err := toml.DecodeFile("configs/networking/networking.toml", &nc); err != nil {
// 		log.Fatalln(err)
// 	}
// 	return nc
// }

// func ChooseBestServer(nc NetworkingConfig) (string, string) {
// 	stat := make(chan AddressWithRtt)
// 	result := make(chan AddressWithRtt)

// 	for _, v := range nc {
// 		go func(addr string, port string) {
// 			faddr, err := creators.CreateAddr(addr, port)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			pinger, err := ping.NewPinger(faddr)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			pinger.Count = 3
// 			if err = pinger.Run(); err != nil {
// 				log.Fatalln(err)
// 			}
// 			stat <- AddressWithRtt{
// 				Address{
// 					Addr: addr,
// 					Port: port,
// 				},
// 				pinger.Statistics().AvgRtt.Microseconds(),
// 			}
// 		}(v.Addr, v.Port)
// 	}
// 	go func() {
// 		var best AddressWithRtt
// 		for s := range stat {
// 			if best.AvgRtt != 0 && best.AvgRtt > s.AvgRtt {
// 				best = s
// 			}
// 		}
// 		result <- best
// 		close(stat)
// 		close(result)
// 	}()
// 	r := <-result
// 	return r.Addr, r.Port
// }

// func GetInstance() *establisher.Establisher {
// 	if conn == nil {
// 		ip, port := ChooseBestServer(GetServersToConnect())
// 		conn = connector.Connect(config.Config{
// 			IP:         ip,
// 			Port:       port,
// 			PingerAddr: "www.google.com",
// 			Workers:    4,
// 		})
// 	}
// 	return conn
// }
