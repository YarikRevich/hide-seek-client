package render

import (
	// metadata "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

// = make(map[*ebiten.Image]func(*ebiten.Image) *ebiten.DrawImageOptions)

var (
	imagesToRender  = make([]RenderCell, 0, 100)
	debugTextToRender = make([]func(*ebiten.Image), 0, 100)
)

type RenderCell struct {
	Image *ebiten.Image
	CallBack func(*ebiten.Image) *ebiten.DrawImageOptions
}

func SetImageToRender(c RenderCell) {
	imagesToRender = append(imagesToRender, c)
}

func GetImagesToRender() []RenderCell {
	return imagesToRender
}

// func GetCollisionsOfImagesToRender()metadata.M{

// }

func SetDebugTextToRender(c func(*ebiten.Image)) {
	debugTextToRender = append(debugTextToRender, c)
}

func GetDebugTextToRender() []func(*ebiten.Image) {
	return debugTextToRender
}
