package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderCallback func(*ebiten.Image)


var (
	// imagesToRender  = make([]Cell, 0, 100)
	// debugTextToRender = make([]func(*ebiten.Image), 0, 100)
	// widgetsToRender = make([]func(*ebiten.Image), 0, 100)
	toRender = make([]RenderCallback, 0, 100)
)

// type Cell struct {
// 	Image *ebiten.Image
// 	CallBack func(*ebiten.Image, *ebiten.Image) *ebiten.DrawImageOptions
// }

func SetToRender(c RenderCallback){
	toRender = append(toRender, c)
}
func GetToRender()[]RenderCallback{
	return toRender
}

// func SetImageToRender(c func(*ebiten.Image)*ebiten.DrawImageOptions) {
// 	imagesToRender = append(imagesToRender, c)
// }

// func GetImagesToRender() []Cell {
// 	return imagesToRender
// }

// func SetWidgetToRender(c func(*ebiten.Image)) {
// 	widgetsToRender = append(widgetsToRender, c)
// }

// func GetWidgetToRender() []func(*ebiten.Image) {
// 	return widgetsToRender
// }

// func SetTextToRender(c func(*ebiten.Image)) {
// 	debugTextToRender = append(debugTextToRender, c)
// }

// func GetTextToRender() []func(*ebiten.Image) {
// 	return debugTextToRender
// }

func CleanRenderPool(){
	toRender = toRender[:0]
 	// debugTextToRender = debugTextToRender[:0]
	// widgetsToRender = widgetsToRender[:0]
}
