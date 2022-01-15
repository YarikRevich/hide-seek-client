package mapchoose

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
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
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("maps/thumbnails/helloween")
		m := sources.UseSources().Metadata().GetMetadata("maps/thumbnails/helloween")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("maps/thumbnails/starwars")
		m := sources.UseSources().Metadata().GetMetadata("maps/thumbnails/starwars")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		screen.DrawImage(img, opts)
	})
}
