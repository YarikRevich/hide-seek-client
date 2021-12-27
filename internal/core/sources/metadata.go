package sources

import (
	"embed"
	"path/filepath"
	"regexp"
	"sync"

	"image"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type TextPosition string
type FontColor string
type GameRole string

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

const (
	GameMap GameRole = "gamemap"
)

type Animation struct {
	FrameDelay,
	FrameX,
	FrameY,
	FrameWidth,
	FrameHeight float64

	FrameNum int
}

//Runtime defined metadata model
type RuntimeDefined struct {
	ZoomedScale struct {
		X, Y float64
	}
}

type Transition struct {
	StartScale, EndScale struct {
		X, Y float64
	}
}

type Model struct {
	RuntimeDefined RuntimeDefined

	Animation Animation

	Transition Transition

	//Information about metadata file
	Info struct {
		//Parent file metadata one related to
		GameRole

		Parent                   string
		ScrollableX, ScrollableY bool
	}

	Size struct {
		Width, Height float64
	}

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
		X, Y float64
	}

	Offset struct {
		X, Y float64
	}

	Text struct {
		Symbols  string
		Position TextPosition
	}

	Fonts struct {
		FontColor
	}

	Camera struct {
		MaxZoom, MinZoom, InitZoom float64
	}

	Effects struct {
		TraceColorBegin [4]float32
		TraceColorEnd   [4]float32
	}
}

type ModelCombination struct {
	Modified, Origin *Model
}

type Metadata struct {
	sync.Mutex

	Collection map[string]*ModelCombination
}

func (m *Metadata) loadFile(fs embed.FS, path string) {
	var o, q Model

	if _, err := toml.DecodeFS(fs, path, &o); err != nil {
		logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
	}
	q = o

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		m.Lock()

		m.Collection[reg.Split(path, -1)[0]] = &ModelCombination{Modified: &q, Origin: &o}
		m.Unlock()
	}
}

func (m *Metadata) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, m.loadFile).Parse()
	wg.Done()
}

func (m *Metadata) GetMetadata(path string) *ModelCombination {
	path = filepath.Join("dist/metadata", path)
	return m.Collection[path]
}

func NewMetadata() *Metadata {
	return &Metadata{Collection: make(map[string]*ModelCombination)}
}
