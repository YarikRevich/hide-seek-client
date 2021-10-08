package debug

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func Draw() {
	// log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------",
	// 	l.userConfig.Pos.X,
	// 	l.userConfig.Pos.Y,
	// 	l.userConfig.PersonalInfo.HeroPicture,
	// 	l.userConfig.PersonalInfo.LobbyID,
	// 	strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0],

	render.SetToRender(func(screen *ebiten.Image) {

		f := fontcollection.GetFontBySize(metadatacollection.GetMetadata("assets/fonts/debug/debug").Fonts.Font)
		text.Draw(screen, fmt.Sprintf(
			"!!!\n%s\n%s\n!!!",
			strconv.FormatFloat(ebiten.CurrentFPS(), 'f', 0, 32),
			strconv.FormatFloat(ebiten.CurrentTPS(), 'f', 0, 32)), f, 0, 15, color.White)
	})
}
