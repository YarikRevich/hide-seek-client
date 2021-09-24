package render

import "github.com/hajimehoshi/ebiten/v2"

var (
	imagesToRender    = make(map[*ebiten.Image]func(*ebiten.Image) *ebiten.DrawImageOptions)
	debugTextToRender = make([]func(*ebiten.Image), 0, 100)
)

func SetImageToRender(i *ebiten.Image, c func(*ebiten.Image) *ebiten.DrawImageOptions) {
	imagesToRender[i] = c
}

func GetImagesToRender() map[*ebiten.Image]func(*ebiten.Image) *ebiten.DrawImageOptions {
	return imagesToRender
}

func SetDebugTextToRender(c func(*ebiten.Image)) {
	debugTextToRender = append(debugTextToRender, c)
}

func GetDebugTextToRender() []func(*ebiten.Image) {
	return debugTextToRender
}
