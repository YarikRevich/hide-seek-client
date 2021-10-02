package settingsmenu

import (
	//"github.com/blizzy78/ebitenui"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {


	render.SetToRender(func(screen *ebiten.Image) {
		img := imageloader.GetImage("assets/images/menues/background/WaitRoom")

		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		screen.DrawImage(img, opts)
	})
}
