package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imagesToRender  = make([]Cell, 0, 100)
	debugTextToRender = make([]func(*ebiten.Image), 0, 100)

)

type Cell struct {
	Image *ebiten.Image
	CallBack func(*ebiten.Image) *ebiten.DrawImageOptions
}

func SetImageToRender(c Cell) {
	imagesToRender = append(imagesToRender, c)
}

func GetImagesToRender() []Cell {
	return imagesToRender
}

func SetTextToRender(c func(*ebiten.Image)) {
	debugTextToRender = append(debugTextToRender, c)
}

func GetTextToRender() []func(*ebiten.Image) {
	return debugTextToRender
}

func CleanRenderPool(){
	imagesToRender = imagesToRender[:0]
 	debugTextToRender = debugTextToRender[:0]
}
