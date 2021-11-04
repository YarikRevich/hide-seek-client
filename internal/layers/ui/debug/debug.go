package debug

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
)

func Draw() {
	// log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------",
	// 	l.userConfig.Pos.X,
	// 	l.userConfig.Pos.Y,
	// 	l.userConfig.PersonalInfo.HeroPicture,
	// 	l.userConfig.PersonalInfo.LobbyID,
	// 	strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0],

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		f := fontcollection.GetFont("assets/fonts/base")

		text.Draw(screen, profiling.UseProfiler().String(), f, 0, 0, color.White)
	})
}
