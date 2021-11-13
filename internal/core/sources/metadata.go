package sources

import (
	"embed"
	"fmt"
	"path/filepath"
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
		CoefficiantX, CoefficiantY float64
	}

	Text struct {
		Symbols  string
		Position TextPosition
	}

	Fonts struct {
		FontColor FontColor
	}

	Camera struct {
		MaxZoom, MinZoom, Zoom float64
	}
}

//Scales margins with scale
func (m *Model) ScaleMargins() (float64, float64) {
	return m.Margins.LeftMargin * m.Scale.CoefficiantX, m.Margins.TopMargin * m.Scale.CoefficiantY
}

type ModelCombination struct {
	Modified, Origin *Model
}

type Metadata struct {
	sync.Mutex

	Used       map[string]*ModelCombination
	Collection map[string]*Model
}

func (m *Metadata) loadFile(fs embed.FS, path string) {
	var o Model

	if _, err := toml.DecodeFS(fs, path, &o); err != nil {
		logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
	}

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		m.Lock()
		m.Collection[reg.Split(path, -1)[0]] = &o
		m.Unlock()
	}
}

func (m *Metadata) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, m.loadFile).Parse()
	wg.Done()
}

func (m *Metadata) GetMetadata(path string) *ModelCombination {
	path = filepath.Join("assets/metadata", path)

	if _, ok := m.Used[path]; !ok {
		file, ok := m.Collection[path]
		if !ok {
			logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
		}

		c := new(ModelCombination)
		q := new(Model)
		*q = *file
		c.Modified = q
		c.Origin = file

		m.Used[path] = c
		return c
	}

	return m.Used[path]
}

func (m *Metadata) CleanUsedMetadata(){
	m.Used = make(map[string]*ModelCombination)
}

func NewMetadata() *Metadata {
	return &Metadata{Collection: make(map[string]*Model),
		Used: make(map[string]*ModelCombination)}
}
