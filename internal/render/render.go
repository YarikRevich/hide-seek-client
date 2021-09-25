package render

import (
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imagesToRender    = make(map[*imageloader.Image]func(*ebiten.Image) *ebiten.DrawImageOptions)
	debugTextToRender = make([]func(*ebiten.Image), 0, 100)
)

func SetImageToRender(i *imageloader.Image, c func(*ebiten.Image) *ebiten.DrawImageOptions) {
	imagesToRender[i] = c
}

func GetImagesToRender() map[*imageloader.Image]func(*ebiten.Image) *ebiten.DrawImageOptions {
	return imagesToRender
}

func SetDebugTextToRender(c func(*ebiten.Image)) {
	debugTextToRender = append(debugTextToRender, c)
}

func GetDebugTextToRender() []func(*ebiten.Image) {
	return debugTextToRender
}
