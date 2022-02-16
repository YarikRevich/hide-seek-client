package world

type CameraOpts struct {
	Zoom     float64
	Angle    float64
	Rotation float64
}

type Camera struct {
	Opts *CameraOpts
}

func (c *Camera) SetOpts(opts *CameraOpts) {
	c.Opts = opts
}
