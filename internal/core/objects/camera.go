package objects

type Camera struct {
	Object

	Zoom float64
}

//Increments zoom property
func (c *Camera) ZoomIn() {
	if c.Zoom < c.ModelCombination.Modified.Camera.MaxZoom {
		c.Zoom++
	}

	// c.Hero.Followed.SetTranslationXMovementBlocked(false)
	// c.Hero.Followed.SetTranslationYMovementBlocked(false)
}

//Decrements zoom property
func (c *Camera) ZoomOut() {
	// wsx, _ := w.GetZoomedMapScale()
	// czx, _ := c.Hero.Followed.GetZoomedRawPosForCamera(w.GetZoomedMapScale())
	// zx, _ := c.Hero.Followed.GetZoomedRawPos(w.GetZoomedMapScale())
	// fmt.Println(zx, czx, m.Size.Width * wsx)
	// co := objects.UseObjects().Camera()
	// fmt.Println(co.RawPos.X, czx)
	// co := objects.UseObjects().Camera()
	// fmt.Println(co.RawPos.X+1 < 1113)
	// if co.RawPos.X+1 < 1113 {
	if c.Zoom > c.ModelCombination.Modified.Camera.MaxZoom {
		c.Zoom--
	}

	// c.Hero.Followed.SetTranslationXMovementBlocked(false)
	// c.Hero.Followed.SetTranslationYMovementBlocked(false)
}

func NewCamera() *Camera {
	return new(Camera)
}
