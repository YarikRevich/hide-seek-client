package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderCallback func(*ebiten.Image)

var toRender = make([]RenderCallback, 0, 100)

func SetToRender(c RenderCallback){
	toRender = append(toRender, c)
}
func GetToRender()[]RenderCallback{
	return toRender
}

func CleanRenderPool(){
	toRender = toRender[:0]
}
