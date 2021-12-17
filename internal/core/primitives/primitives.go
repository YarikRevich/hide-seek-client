package primitives

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//Create square by ebiten tools
func CreateSquare(a int) *ebiten.Image {
	return ebiten.NewImage(a, a)
}
