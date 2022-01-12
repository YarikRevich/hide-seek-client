package objects

import (
	"math"

	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
)

type Camera struct {
	Base

	Zoom        float64
	AlignOffset struct {
		X, Y float64
	}
}

func (c *Camera) GetZoom() sources.Vec2 {
	s := c.MetadataModel.GetScale()
	return sources.Vec2{X: s.X / 100 * c.Zoom, Y: s.Y / 100 * c.Zoom}
}

//Increments zoom property
func (c *Camera) ZoomIn(o *Base) {
	if c.Zoom < c.MetadataModel.Camera.MaxZoom {
		c.AlignOffset.X += 11
		c.AlignOffset.Y += 9
		c.Zoom++
		// o.SetTranslationYMovementBlocked(false)
		// o.SetTranslationXMovementBlocked(false)
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut(o *Base) {
	if !math.Signbit(c.GetScaledPosX()+c.AlignOffset.X) && !math.Signbit(c.GetScaledPosY()+c.AlignOffset.Y) {
		if c.Zoom > c.MetadataModel.Camera.MinZoom {
			c.AlignOffset.X -= 11
			c.AlignOffset.Y -= 9
			c.Zoom--
			// o.SetTranslationYMovementBlocked(false)
			// o.SetTranslationXMovementBlocked(false)
		}
	}
}

func NewCamera() *Camera {
	c := new(Camera)
	c.SetSkin("camera/camera")
	c.Zoom = c.MetadataModel.Camera.InitZoom
	return c
}
