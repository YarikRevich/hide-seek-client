package settingsmenu

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text/positioning"

	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/background/background")

		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

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

	f := sources.UseSources().Font().GetFont("base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/inputs/input")
		m := sources.UseSources().Metadata().GetMetadata("system/inputs/input").Modified

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		t := events.UseEvents().Input().SettingsMenuNameBuffer.Read()

		s := positioning.UsePositioning().Input()
		s.Init(img, m, f, t)
		s.Draw()

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/button")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/button").Modified

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		s := positioning.UsePositioning().Input()
		s.Init(img, m, f, m.Text.Symbols)
		s.Draw()

		screen.DrawImage(img, opts)
	})
}
