package objects

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type Camera struct {
	Base

	Zoom    float64
	COffset types.Vec2
	CScale  types.Vec2

	// fOffsetX += (fMouseWorldX_BeforeZoom - fMouseWorldX_AfterZoom);
	// 	fOffsetY += (fMouseWorldY_BeforeZoom - fMouseWorldY_AfterZoom);
}

func (c *Camera) GetZoom() types.Vec2 {
	s := c.MetadataModel.GetScale()
	return types.Vec2{X: s.X / 100 * c.Zoom, Y: s.Y / 100 * c.Zoom}
}

//Increments zoom property
func (c *Camera) ZoomIn(o *Base) {
	// pScaleOld := c.GetZoomedScale(o)

	cPosOld := c.ConvertScreenToWorld(c.RawPos)
	if c.Zoom < c.MetadataModel.Camera.MaxZoom {
		// c.Zoom = 2.001
		// c.Zoom++
		c.CScale.X += 0.05
		c.CScale.Y += 0.05

		// pScale := c.GetZoomedScale(o)
		cPos := c.ConvertScreenToWorld(c.RawPos)
		c.COffset.X += cPosOld.X - cPos.X
		c.COffset.Y += cPosOld.Y - cPos.Y
		// o.SetTranslationYMovementBlocked(false)
		// o.SetTranslationXMovementBlocked(false)
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut(o *Base) {
	// pScaleOld := c.GetZoomedScale(o)
	cPosOld := c.ConvertScreenToWorld(c.RawPos)
	if c.Zoom > c.MetadataModel.Camera.MinZoom {
		// c.AlignOffset.X -= pScale.X
		// c.AlignOffset.Y += pScale.Y
		// c.Zoom--
		c.CScale.X -= 0.05
		c.CScale.Y -= 0.05

		cPos := c.ConvertScreenToWorld(c.RawPos)
		c.COffset.X += cPosOld.X - cPos.X
		c.COffset.Y += cPosOld.Y - cPos.Y
		// fmt.Println(cPos, cPosOld)
		// o.SetTranslationYMovementBlocked(false)
		// o.SetTranslationXMovementBlocked(false)
	}
	// }
	// }
}

func (c *Camera) GetZoomedScale(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	return types.Vec2{X: ((s.X * c.Zoom) / 100), Y: ((s.Y * c.Zoom) / 100)}
}

func (c *Camera) GetZoomedPos(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	return types.Vec2{X: b.RawPos.X * ((s.X * c.Zoom) / 100), Y: b.RawPos.Y * ((s.Y * c.Zoom) / 100)}
}

func (c *Camera) GetZoomedOffset(b *Base) types.Vec2 {
	s := b.MetadataModel.GetScale()
	return types.Vec2{X: b.RawOffset.X * ((s.X * c.Zoom) / 100), Y: b.RawOffset.Y * ((s.Y * c.Zoom) / 100)}
}

//Checks if passed v2 is outta range of v1
func (c *Camera) IsOuttaRange(v1, v2 float64) bool {
	return v2 > v1
}

func (c *Camera) ConvertWorldToScreen(w types.Vec2) types.Vec2 {
	return types.Vec2{X: (w.X - c.COffset.X) * c.Scale.X, Y: (w.Y - c.COffset.Y) * c.Scale.Y}
}
func (c *Camera) ConvertScreenToWorld(s types.Vec2) types.Vec2 {
	return types.Vec2{X: (s.X / c.Scale.X) + c.COffset.X, Y: (s.Y / c.Scale.Y) + c.COffset.Y}
}

func NewCamera() *Camera {
	c := new(Camera)
	c.SetSkin("camera/camera")
	c.Zoom = c.MetadataModel.Camera.InitZoom
	c.Type = "camera"
	c.COffset.X = -screen.UseScreen().GetAxis().X
	c.COffset.Y = -screen.UseScreen().GetAxis().Y
	c.CScale.X = 1
	c.CScale.Y = 1
	return c
}
