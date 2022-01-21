package objects

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type Camera struct {
	Base

	Zoom        float64
	AlignOffset types.Vec2

	DeltaOffset types.Vec2
}

func (c *Camera) GetZoom() types.Vec2 {
	s := c.MetadataModel.GetScale()
	return types.Vec2{X: s.X / 100 * c.Zoom, Y: s.Y / 100 * c.Zoom}
}

//Increments zoom property
func (c *Camera) ZoomIn(o *Base) {
	pScale := c.GetZoomedScale(o)
	cPosOld := types.Vec2{X: c.RawPos.X * pScale.X, Y: c.RawPos.Y * pScale.Y}
	if c.Zoom < c.MetadataModel.Camera.MaxZoom {
		// c.AlignOffset.X += pScale.X
		// c.AlignOffset.Y -= pScale.Y

		c.Zoom++
		pScale = c.GetZoomedScale(o)
		cPos := types.Vec2{X: c.RawPos.X * pScale.X, Y: c.RawPos.Y * pScale.Y}
		c.DeltaOffset.X += cPos.X - cPosOld.X
		c.DeltaOffset.Y += cPos.Y - cPosOld.Y
		// o.SetTranslationYMovementBlocked(false)
		// o.SetTranslationXMovementBlocked(false)
	}
}

//Decrements zoom property
func (c *Camera) ZoomOut(o *Base) {
	pScaleOld := c.GetZoomedScale(o)
	cPosOld := types.Vec2{X: c.RawPos.X * pScaleOld.X, Y: c.RawPos.Y * pScaleOld.Y}
	// fmt.Println(!math.Signbit(cPos.X+c.AlignOffset.X), !math.Signbit(cPos.Y+c.AlignOffset.Y))
	// if !math.Signbit(cPos.X+c.AlignOffset.X) && !math.Signbit(cPos.Y+c.AlignOffset.Y) {
	// if !math.Signbit(c.GetScaledPosX()+c.AlignOffset.X) && !math.Signbit(c.GetScaledPosY()+c.AlignOffset.Y) {
	if c.Zoom > c.MetadataModel.Camera.MinZoom {
		// c.AlignOffset.X -= pScale.X
		// c.AlignOffset.Y += pScale.Y
		c.Zoom--
		pScale := c.GetZoomedScale(o)
		cPos := types.Vec2{X: c.RawPos.X * pScale.X, Y: c.RawPos.Y * pScale.Y}
		c.DeltaOffset.X += cPos.X - cPosOld.X
		c.DeltaOffset.Y += cPos.Y - cPosOld.Y

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

func NewCamera() *Camera {
	c := new(Camera)
	c.SetSkin("camera/camera")
	c.Zoom = c.MetadataModel.Camera.InitZoom
	c.Type = "camera"
	return c
}
