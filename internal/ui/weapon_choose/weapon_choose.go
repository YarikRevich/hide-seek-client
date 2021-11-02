package weaponchoose

import (
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/hajimehoshi/ebiten/v2"
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

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	img := imagecollection.GetImage("assets/images/maps/thumbnails/helloween")
	// 	m := metadatacollection.GetMetadata("assets/images/maps/thumbnails/helloween")

	// 	opts := &ebiten.DrawImageOptions{}

	// 	opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
	// 	opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

	// 	screen.DrawImage(img, opts)
	// })

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	img := imagecollection.GetImage("assets/images/maps/thumbnails/starwars")
	// 	m := metadatacollection.GetMetadata("assets/images/maps/thumbnails/starwars")

	// 	opts := &ebiten.DrawImageOptions{}

	// 	opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
	// 	opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

	// 	screen.DrawImage(img, opts)
	// })
}
