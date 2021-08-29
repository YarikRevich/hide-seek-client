package debug

import (
	"log"
	"os"
	"strings"
)

func DrawDebugInfo() {
	log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------",
		l.userConfig.Pos.X, 
		l.userConfig.Pos.Y, 
		l.userConfig.PersonalInfo.HeroPicture, 
		l.userConfig.PersonalInfo.LobbyID, 
		strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0], 
		l.stat.AvgRtt, 
		l.stat.PacketLoss, 
		l.stat.PacketsSent, 
		l.stat.PacketsRecv,
	)
}
