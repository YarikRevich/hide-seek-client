package waitroomjoin

import (
	"fmt"
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
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

		worldMap := world.UseWorld().GetWorldMap()
		text.Draw(img, fmt.Sprintf("World ID: %s", worldMap.ID), f, 340, 60, color.White)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/back")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/textareas/textarea")
		m := sources.UseSources().Metadata().GetMetadata("system/textareas/textarea")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		text.Draw(img, world.UseWorld().String(), f, 10, 20, &color.RGBA{100, 100, 100, 255})

		screen.DrawImage(img, opts)
	})
}
