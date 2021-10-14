package wait_room

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/interface/positioning/button"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/background/background")

		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/buttons/back")
		m := metadatacollection.GetMetadata("assets/images/system/buttons/back")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/textareas/textarea")
		m := metadatacollection.GetMetadata("assets/images/system/textareas/textarea")

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		// fmt.Println(m.Fonts.Font)
		f := fontcollection.GetFontBySize(m.Fonts.Font)
		text.Draw(img, world.UseWorld().FormatUsersUsername(), f, 10, 25, &color.RGBA{100, 100, 100, 255})

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_confirm_game")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fontcollection.GetFontBySize(m.Fonts.Font)
		tx, ty := button.ChooseButtonTextPosition(f, m.Button.Text, *m)
		
		text.Draw(img, m.Button.Text, f, tx, ty, color.White)

		screen.DrawImage(img, opts)
	})
	// l.winConf.TextAreas.NewMembersAnnouncement.Clear()
	// l.winConf.TextAreas.NewMembersAnnouncement.Write([]byte("Wait for members!"))
	// l.winConf.TextAreas.NewMembersAnnouncement.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.NewMembersAnnouncement.Orig, 4))

	// l.winConf.TextAreas.NewMembersTextArea.Clear()
	// for _, value := range l.winConf.WaitRoom.NewMembers{
	// 	l.winConf.TextAreas.NewMembersTextArea.Write([]byte(value + "\n"))
	// }
	// l.winConf.TextAreas.NewMembersTextArea.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.NewMembersTextArea.Orig, 2.5))

	// l.winConf.TextAreas.CurrentLobbyIDArea.Clear()
	// lobbyIdText := fmt.Sprintf("Lobby ID is: %s", l.userConfig.PersonalInfo.LobbyID)
	// l.winConf.TextAreas.CurrentLobbyIDArea.Write([]byte(lobbyIdText))
	// l.winConf.TextAreas.CurrentLobbyIDArea.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.CurrentLobbyIDArea.Orig, 2.5))

}
