package world

import (
	"math"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/kvartborg/vector"
)

type Camera struct {
	Position           types.Vec3
	Zoom, Angle, Pitch float64

	Rotation types.Matrix4
}

func (c *Camera) ZoomIn(v float64) {
	c.Zoom += v
}

func (c *Camera) ZoomOut(v float64) {
	c.Zoom -= v
}

func (c *Camera) MoveAngle(v float64) {
	c.Angle += v
}

func (c *Camera) MovePitch(v float64) {
	c.Pitch += v
}

func (c *Camera) MovePositionX(v float64) {
	c.Position.X += v
}

func (c *Camera) MovePositionY(v float64) {
	c.Position.Y += v
}

func (c *Camera) MovePositionZ(v float64) {
	c.Position.Z += v
}

func (ca *Camera) Rotate(x, y, z, angle float64) {
	mat := types.Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	vector := vector.Vector{x, y, z}.Unit()
	s := math.Sin(angle)
	c := math.Cos(angle)
	m := 1 - c

	mat[0][0] = m*vector[0]*vector[0] + c
	mat[0][1] = m*vector[0]*vector[1] + vector[2]*s
	mat[0][2] = m*vector[2]*vector[0] - vector[1]*s

	mat[1][0] = m*vector[0]*vector[1] - vector[2]*s
	mat[1][1] = m*vector[1]*vector[1] + c
	mat[1][2] = m*vector[1]*vector[2] + vector[0]*s

	mat[2][0] = m*vector[2]*vector[0] + vector[1]*s
	mat[2][1] = m*vector[1]*vector[2] - vector[0]*s
	mat[2][2] = m*vector[2]*vector[2] + c

	ca.Rotation = mat.GetMultipied(ca.Rotation)
}

func (c *Camera) GetProjection(sm *screen.ScreenManager, near, far, right, left, top, bottom float64) types.Matrix4 {
	w, h := sm.Image.Size()
	asr := float64(h) / float64(w)

	return types.Matrix4{
		{2 / (1*c.Zoom - (-1 * c.Zoom)), 0, 0, 0},
		{0, 2 / (asr*c.Zoom - (-asr * c.Zoom)), 0, 0},
		{0, 0, -2, 0},
		{0, 0, 0, 1},
	}
}

func (camera *Camera) GetView() types.Matrix4 {
	var mat types.Matrix4
	mat[3][0] = camera.Position.X
	mat[3][1] = camera.Position.Y
	mat[3][2] = camera.Position.Z

	return mat.GetMultipied(camera.Rotation.GetTransposed())
}
