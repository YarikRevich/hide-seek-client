package world

import "github.com/YarikRevich/hide-seek-client/internal/core/types"

type Camera struct {
	Position           types.Vec3
	Zoom, Angle, Pitch float64
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
