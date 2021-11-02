package wait_room

import (
	"fmt"
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	f := fontcollection.GetFont("assets/fonts/base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/background/background"))

		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))


		m := metadatacollection.GetMetadata("assets/fonts/waitroom/waitroom")
		w := objects.UseObjects().World()

		ebitentext.Draw(img, fmt.Sprintf("World ID: %s", w.ID), f, int(m.Margins.LeftMargin), int(m.Margins.TopMargin), color.White)

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
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/textareas/textarea"))
		m := metadatacollection.GetMetadata("assets/images/system/textareas/textarea")

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		ebitentext.Draw(img, objects.UseObjects().World().String(), f, 10, 20, &color.RGBA{100, 100, 100, 255})

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {

	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_confirm_game")

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

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
