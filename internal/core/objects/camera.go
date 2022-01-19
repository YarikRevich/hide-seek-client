package objects

import "github.com/YarikRevich/hide-seek-client/internal/core/types"

type Camera struct {
	Base

	Zoom        float64
	AlignOffset types.Vec2
}

func (c *Camera) GetZoom() types.Vec2 {
	s := c.MetadataModel.GetScale()
	return types.Vec2{X: s.X / 100 * c.Zoom, Y: s.Y / 100 * c.Zoom}
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
	// if !math.Signbit(c.GetScaledPosX()+c.AlignOffset.X) && !math.Signbit(c.GetScaledPosY()+c.AlignOffset.Y) {
	if c.Zoom > c.MetadataModel.Camera.MinZoom {
		c.AlignOffset.X -= 11
		c.AlignOffset.Y -= 9
		c.Zoom--
		// o.SetTranslationYMovementBlocked(false)
		// o.SetTranslationXMovementBlocked(false)
	}
	// }
}

//Returns zoomed offset for passed object
func (c *Camera) GetZoomedOffset(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	o := c.MetadataModel.GetOffset()
	return types.Vec2{X: (b.RawOffset.X * (s.X * c.Zoom / 100)) - o.X, Y: (b.RawOffset.Y * (s.Y * c.Zoom / 100)) - o.Y}
}

//Returns zoomed position for passed object
func (c *Camera) GetZoomedPos(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	o := c.MetadataModel.GetOffset()
	return types.Vec2{X: (b.RawPos.X * (s.X * c.Zoom / 100)) - o.X, Y: (b.RawPos.Y * (s.Y * c.Zoom / 100)) - o.Y}
}

func (c *Camera) GetZoomedScale(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	return types.Vec2{X: (s.X * c.Zoom / 100), Y: (s.Y * c.Zoom / 100)}
}

func NewCamera() *Camera {
	c := new(Camera)
	c.SetSkin("camera/camera")
	c.Zoom = c.MetadataModel.Camera.InitZoom
	c.Type = "camera"
	return c
}
