package settingsmenu

import (
	"fmt"

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
		img := sources.UseSources().Images().GetCopyOfImage("system/inputs/input")
		mm := sources.UseSources().Metadata().GetMetadata("system/inputs/input").Modified
		mo := sources.UseSources().Metadata().GetMetadata("system/inputs/input").Origin

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(mm.Margins.LeftMargin, mm.Margins.TopMargin)
		opts.GeoM.Scale(mm.Scale.CoefficiantX, mm.Scale.CoefficiantY)

		t := events.UseEvents().Input().SettingsMenuNameBuffer.Read()
		fmt.Println(t)

		s := positioning.UsePositioning().Input()
		s.Init(img, mo, f, t)
		s.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/buttons/button")
		mm := sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config").Modified
		mo := sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config").Origin

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(mm.Margins.LeftMargin, mm.Margins.TopMargin)
		opts.GeoM.Scale(mm.Scale.CoefficiantX, mm.Scale.CoefficiantY)

		s := positioning.UsePositioning().Button()
		s.Init(img, mo, f, mo.Text.Symbols)
		s.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})
}