package networking

import (
	"github.com/YarikRevich/HideSeek-Client/internal/networking/game"
	joinlobbymenu "github.com/YarikRevich/HideSeek-Client/internal/networking/join_lobby_menu"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/state_machine/ui"

)

func Process(){
	switch ui.UseStatus().GetState(){
	case ui.JOIN_LOBBY_MENU:
		joinlobbymenu.Exec()
	case ui.GAME:
		game.Exec()
	}
}





// func connectToMainServer(adress string)*net.UDPConn{

// 	splittedAdress := strings.Split(strings.Replace(adress, "\n", "", 1), ":")
// 	ip, portStr := splittedAdress[0], splittedAdress[1]
// 	port, err := strconv.Atoi(portStr)
// 	if err != nil{
// 		panic(err)
// 	}
// 	udpaddr := net.UDPAddr{
// 		Port: port,
// 		IP: net.ParseIP(ip),
// 	}
// 	resolvedadd, err := net.ResolveUDPAddr("udp", udpaddr.String())

// 	conn, err := net.DialUDP("udp", nil, resolvedadd)
// 	if err != nil{
// 		color.Red.Println("There is no connection for the internet")
// 		os.Exit(0)
// 	}
// 	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
// 	conn.SetWriteDeadline(time.Now().Add(100 * time.Millisecond))
// 	return conn
// }

// func getAvailableServers(conn *net.UDPConn)map[int]string{
// 	defer conn.Close()
// 	parser := Server.StartParser(new(Server.StartRequest))

// 	server := Server.Network(new(Server.N))
// 	server.Init(conn, nil, 1, parser.Parse, nil, "CheckServers")
// 	server.Write()

// 	data := server.ReadStart(parser.Unparse)
// 	result := make(map[int]string)
// 	for index, value := range data[0].Body{
// 		result[index+1] = value
// 	}
// 	return result
// }

// func formatAvailableServersList(availableservers map[int]string)error{

// 	value, _ := availableservers[1]
// 	if len(availableservers) == 1 && len(value) == 0{
// 		color.Red.Println("There are no available servers right now!")
// 		return errors.New("there are no available server")
// 	}
// 	for index, value := range availableservers{
// 		ip := strings.Split(value, ":")[0]
// 		pink := color.Magenta.Darken().Render
// 		fmt.Printf("%s %s: %d)\n", pink("=>"), ip, index)
// 	}
// 	return nil
// }