package settingsmenu

import (
	"image/color"

	// buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text"
	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"

	// "github.com/YarikRevich/HideSeek-Client/internal/interface/positioning"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
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
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/inputs/input"))
		m := metadatacollection.GetMetadata("assets/images/system/inputs/input")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fontcollection.GetFont("assets/fonts/base")
		
		
		t := events.UseEvents().Input().SettingsMenuNameBuffer.Read()
		
		p := text.NewPositionSession(
			f, 
			t, 
			m.RawSize.Width,
			m.RawSize.Height, 
			m.Button.TextPosition)

		for p.Next() {
			tx, ty := p.GetPosition()
			ebitentext.Draw(
				img,
				p.GetText(),
				f,
				tx, ty,
				color.White)
		}

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_save_config")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fontcollection.GetFont("assets/fonts/base")
		p := text.NewPositionSession(
			f, 
			m.Button.Text, 
			m.RawSize.Width,
			m.RawSize.Height, 
			m.Button.TextPosition)

		for p.Next() {
			tx, ty := p.GetPosition()
			ebitentext.Draw(
				img,
				p.GetText(),
				f,
				tx, ty,
				color.White)
		}

		screen.DrawImage(img, opts)
	})
}
