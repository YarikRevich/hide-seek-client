package waitroom

import (
	"fmt"
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text/positioning"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	f := sources.UseSources().Font().GetFont("base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		
		img := sources.UseSources().Images().GetImage("system/background/background")

		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))


		m := sources.UseSources().Metadata().GetMetadata("waitroom/waitroom").Modified
		w := objects.UseObjects().World()

		text.Draw(img, fmt.Sprintf("World ID: %s", w.ID), f, int(m.Margins.LeftMargin), int(m.Margins.TopMargin), color.White)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/back")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		screen.DrawImage(img, opts)
	})

	

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/textareas/textarea")
		m := sources.UseSources().Metadata().GetMetadata("system/textareas/textarea").Modified

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		text.Draw(img, objects.UseObjects().World().String(), f, 10, 20, &color.RGBA{100, 100, 100, 255})

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {

	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/button")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		s := positioning.UsePositioning().Input()
		s.Init(img, m, f, m.Text.Symbols)
		s.Draw()

		screen.DrawImage(img, opts)
	})
}
