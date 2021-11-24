package waitroomstart

import (
	"fmt"
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text/positioning"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	f := sources.UseSources().Font().GetFont("base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		
		img := sources.UseSources().Images().GetCopyOfImage("system/background/background")

		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))


		mo := sources.UseSources().Metadata().GetMetadata("fonts/waitroom/waitroom").Origin
		worldMap := world.UseWorld().GetWorldMap()

		text.Draw(img, fmt.Sprintf("World ID: %s", worldMap.ID), f, int(mo.Margins.LeftMargin), int(mo.Margins.TopMargin), color.White)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/back")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

		screen.DrawImage(img, opts)
	})

	

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/textareas/textarea")
		m := sources.UseSources().Metadata().GetMetadata("system/textareas/textarea").Modified

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

		
		text.Draw(img, world.UseWorld().String(), f, 10, 20, &color.RGBA{100, 100, 100, 255})

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/buttons/button")
		mm := sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Modified
		mo := sources.UseSources().Metadata().GetMetadata("system/buttons/button_confirm_game").Origin

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(mm.Margins.LeftMargin, mm.Margins.TopMargin)
		opts.GeoM.Scale(mm.Scale.X, mm.Scale.Y)
	
		s := positioning.UsePositioning().Button()
		s.Init(img, mo, f, mo.Text.Symbols)
		s.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})
}