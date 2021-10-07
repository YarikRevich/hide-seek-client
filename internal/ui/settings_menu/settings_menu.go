package settingsmenu

import (
	"image/color"

	buffercollection "github.com/YarikRevich/HideSeek-Client/internal/hid/keyboard/buffers/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/interface/fonts"

	"github.com/YarikRevich/HideSeek-Client/internal/interface/positioning/button"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	render.SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/background/background")

		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		screen.DrawImage(img, opts)
	})

	render.SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/buttons/back")
		m := metadatacollection.GetMetadata("assets/images/system/buttons/back")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		screen.DrawImage(img, opts)
	})

	render.SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/inputs/input"))
		m := metadatacollection.GetMetadata("assets/images/system/inputs/input")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fonts.GetFont(*m)
		t := buffercollection.SettingsMenuNameBuffer.Read()
		tx, ty := button.ChooseButtonTextPosition(f, t, *m)

		text.Draw(img, t, f, tx, ty, color.White)

		screen.DrawImage(img, opts)
	})

	render.SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_save_config")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fonts.GetFont(*m)
		tx, ty := button.ChooseButtonTextPosition(f, m.Button.Text, *m)

		text.Draw(img, m.Button.Text, f, tx, ty,color.White)

		screen.DrawImage(img, opts)
	})
}
