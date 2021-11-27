package objects

import (
	"math"
)

type Camera struct {
	Base

	Zoom        float64
	AlignOffset struct {
		X, Y float64
	}
}

//Increments zoom property
func (c *Camera) ZoomIn(o *Base) {
	if c.Zoom < c.ModelCombination.Modified.Camera.MaxZoom {
		c.AlignOffset.X += 10
		c.AlignOffset.Y += 10
		c.Zoom++
		o.SetTranslationYMovementBlocked(false)
		o.SetTranslationXMovementBlocked(false)
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut(o *Base) {
	if !math.Signbit(c.GetScaledPosX()+c.AlignOffset.X) && !math.Signbit(c.GetScaledPosY()+c.AlignOffset.Y) {
		if c.Zoom > c.ModelCombination.Modified.Camera.MinZoom {
			c.AlignOffset.X -= 10
			c.AlignOffset.Y -= 10
			c.Zoom--
			o.SetTranslationYMovementBlocked(false)
			o.SetTranslationXMovementBlocked(false)
		}
	}
}

func NewCamera() *Camera {
	c := new(Camera)
	c.SetSkin("camera/camera")
	c.Zoom = c.Modified.Camera.InitZoom
	return c
}
