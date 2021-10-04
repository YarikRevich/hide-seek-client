package models

type Position string

const (
	Center Position = "center"
	Left Position = "left"
	Right Position = "right"
)

type Metadata struct {
	Size struct {
		Width  float64
		Height float64
	}
	Margins struct {
		LeftMargin float64
		TopMargin  float64
	}
	Animation struct {
		Delay       float64
		FrameNum    float64
		FrameX      float64
		FrameY      float64
		FrameWidth  float64
		FrameHeight float64
	}
	Spawns []struct {
		X float64
		Y float64
	}
	Physics struct {
		G float64
	}
	Scale struct {
		CoefficiantX float64
		CoefficiantY float64
	}
	Button struct {
		Text string
		Position Position
	}
}
