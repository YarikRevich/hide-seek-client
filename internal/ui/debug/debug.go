package debug

import (
	"fmt"
	"strconv"

	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw() {
	// log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------",
	// 	l.userConfig.Pos.X,
	// 	l.userConfig.Pos.Y,
	// 	l.userConfig.PersonalInfo.HeroPicture,
	// 	l.userConfig.PersonalInfo.LobbyID,
	// 	strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0],
	if cli.GetDebug() {
		render.SetTextToRender(func(screen *ebiten.Image) {
			ebitenutil.DebugPrint(screen,
				fmt.Sprintf(
					"%s\n%s",
					strconv.FormatFloat(ebiten.CurrentFPS(), 'f', 0, 32),
					"text"))
		})
	}
}
