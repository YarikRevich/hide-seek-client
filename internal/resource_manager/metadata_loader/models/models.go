package models

import "image"

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
	FrameDelay,
	FrameNum,
	FrameX,
	FrameY,
	FrameWidth,
	FrameHeight float64
}

type Metadata struct {
	Animation Animation

	//Information about metadata file
	Info struct {
		//Parent file metadata one related to
		Parent     string
		ScrollableX, ScrollableY bool
	}

	//HIDDEN: should not be defined by user by configuration
	Size struct {
		Width, Height float64
	}

	//MUSTN'T be changed over the game
	RawSize struct {
		Width, Height float64
	}

	Margins struct {
		LeftMargin, TopMargin float64
	}

	//MUSTN'T be changed over the game
	RawMargins struct {
		LeftMargin, TopMargin float64
	}

	Spawns []image.Point

	Physics struct {
		G float64
	}

	Buffs struct {
		//MUSTN"T be changed over the game
		RawSpeed struct {
			X, Y float64
		}

		Speed struct {
			X, Y float64
		}
	}

	Scale struct {
		CoefficiantX, CoefficiantY float64
	}

	//MUSN'T be changed over the game
	RawScale struct {
		CoefficiantX, CoefficiantY float64
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
func (m *Metadata) FastenMarginsWithCoefficients() (float64, float64) {
	return m.Margins.LeftMargin * m.Scale.CoefficiantX, m.Margins.TopMargin * m.Scale.CoefficiantY
}
