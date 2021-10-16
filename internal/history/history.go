package history

import (
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenSize struct {
	Height int
	Width  int
}

var (
	lastDirection  direction.Direction
	lastScreenSize ScreenSize
)

func SetDirection(d direction.Direction) {
	lastDirection = d
}

func GetDirection() direction.Direction {
	return lastDirection
}

func UpdateScreenSize(screen *ebiten.Image) {
	w, h := screen.Size()
	lastScreenSize = ScreenSize{
		Height: h, Width: w}
}

func GetScreenSize() (int, int) {
	return lastScreenSize.Width, lastScreenSize.Height
}
