package models

type TextPosition string
type FontColor string

//Predefined options of allowed text positions
const (
	Center TextPosition = "center"
	Left   TextPosition = "left"
	Right  TextPosition = "right"
)

const (
	White FontColor = "white"
	Black FontColor = "black"
)

type Animation struct {
	Delay       float64
	FrameNum    float64
	FrameX      float64
	FrameY      float64
	FrameWidth  float64
	FrameHeight float64
}



type Metadata struct {
	Animation Animation

	//Information about metadata file
	Info struct {
		//Parent file metadata one related to
		Parent string
	}

	//HIDDEN: should not be defined by user by configuration
	Size struct {
		Width  float64
		Height float64
	}

	//MUSTN'T be changed over the project
	RawSize struct {
		Width  float64
		Height float64
	}

	Margins struct {
		LeftMargin float64
		TopMargin  float64
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
		Text         string
		TextPosition TextPosition
	}

	Fonts struct {
		FontColor FontColor
	}
}

//Multiples margins by related coefficients
func (m *Metadata) FastenMarginsWithCoefficients()(float64, float64){
	return m.Margins.LeftMargin * m.Scale.CoefficiantX, m.Margins.TopMargin * m.Scale.CoefficiantY
}
