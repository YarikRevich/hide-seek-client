package debug

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/profiling"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	// log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------",
	// 	l.userConfig.Pos.X,
	// 	l.userConfig.Pos.Y,
	// 	l.userConfig.PersonalInfo.HeroPicture,
	// 	l.userConfig.PersonalInfo.LobbyID,
	// 	strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0],

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		f := sources.UseSources().Font().GetFont("assets/fonts/base")

		text.Draw(screen, profiling.UseProfiler().String(), f, 0, 0, color.White)
	})
}
