package settingsmenu

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(){
	back := imageloader.Images[imageloader.PathsToHash["/images/menues/background/WaitRoom"]]

	fmt.Println("IT WORKS")
	render.SetImageToRender(render.RenderCell{Image: back, CallBack: func(screen *ebiten.Image)*ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := back.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		return opts
	}})

}