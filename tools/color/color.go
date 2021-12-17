package color

import (
	"image/color"

	"github.com/engoengine/glm"
)

//Just creates quat color from array
func CreateColorFromArray(a [4]float32) glm.Quat {
	return glm.Quat{V: glm.Vec3{a[0], a[1], a[2]}, W: a[3]}
}

//Just create rgba color using quat
func CreateRGBAFromQuatColor(q glm.Quat) color.Color {
	return color.RGBA{R: uint8(q.X()), G: uint8(q.Y()), B: uint8(q.Z()), A: uint8(q.W)}
}
