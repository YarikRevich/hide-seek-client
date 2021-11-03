package sources

import (
	"embed"
	"fmt"
	"regexp"
	"sync"

	"image"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

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

type Model struct {
	Animation Animation

	//Information about metadata file
	Info struct {
		//Parent file metadata one related to
		Parent                   string
		ScrollableX, ScrollableY bool
	}

	// //HIDDEN: should not be defined by user by configuration
	// Size struct {
	// 	Width, Height float64
	// }

	//MUSTN'T be changed over the game
	Size struct {
		Width, Height float64
	}

	// Margins struct {
	// 	LeftMargin, TopMargin float64
	// }

	//MUSTN'T be changed over the game
	Margins struct {
		LeftMargin, TopMargin float64
	}

	Spawns []image.Point

	Physics struct {
		G float64
	}

	Buffs struct {
		Speed struct {
			X, Y float64
		}
	}

	Scale struct {
		CoefficiantX, CoefficiantY float64
	}

	//MUSN'T be changed over the game
	// RawScale struct {
	// 	CoefficiantX, CoefficiantY float64
	// }

	Button struct {
		Text         string
		TextPosition TextPosition
	}

	Fonts struct {
		FontColor FontColor
	}
}

//Multiples margins by related coefficients
func (m *Model) FastenMarginsWithCoefficients() (float64, float64) {
	return m.Margins.LeftMargin * m.Scale.CoefficiantX, m.Margins.TopMargin * m.Scale.CoefficiantY
}

type Metadata struct {
	sync.Mutex

	collection map[string]*Model
}

func (m *Metadata) loadFile(fs embed.FS, path string) {
	o := new(Model)

	if _, err := toml.DecodeFS(fs, path, o); err != nil {
		logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
	}

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		m.Lock()
		m.collection[reg.Split(path, -1)[0]] = o
		m.Unlock()
	}
}

func (m *Metadata) Load(fs embed.FS, path string) {
	NewParser(fs, path, m.loadFile).Parse()
}

func (m *Metadata) GetMetadata(path string) *Model {
	file, ok := m.collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return file
}

func NewMetadata() *Metadata {
	return new(Metadata)
}
