package herochoose

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
		img := sources.UseSources().Images().GetImage("heroes/thumbnails/pumpkin")
		m := sources.UseSources().Metadata().GetMetadata("heroes/thumbnails/pumpkin")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Scale(s.X, s.Y)
		opts.GeoM.Translate(ms.X, ms.Y)

		screen.DrawImage(img, opts)
	})

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	img := imagecollection.GetImage("assets/images/maps/thumbnails/helloween")
	// 	m := metadatacollection.GetMetadata("assets/images/maps/thumbnails/helloween")

	// 	opts := &ebiten.DrawImageOptions{}

	// 	opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
	// 	opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

	// 	screen.DrawImage(img, opts)
	// })

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	img := imagecollection.GetImage("assets/images/maps/thumbnails/starwars")
	// 	m := metadatacollection.GetMetadata("assets/images/maps/thumbnails/starwars")

	// 	opts := &ebiten.DrawImageOptions{}

	// 	opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
	// 	opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

	// 	screen.DrawImage(img, opts)
	// })
}
