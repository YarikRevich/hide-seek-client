package sources

import (
	"embed"
	"fmt"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/tools/utils"
	"github.com/sirupsen/logrus"
)

type TextPosition string
type FontColor string
type GameRole string
type Type []string

func (t *Type) Contains(q string) bool {
	for _, v := range *t {
		if v == q {
			return true
		}
	}
	return false
}

type Vec2 struct{ X, Y float64 }

//Predefined options of allowed text positions
const (
	Center TextPosition = "center"
	Left                = "left"
	Right               = "right"
)

const (
	White FontColor = "white"
	Black           = "black"
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
	ZoomedScale Vec2
}

type Transition struct {
	StartScale, EndScale Vec2
}

type MetadataModel struct {
	Type

	Animation Animation

	Transition Transition

	//Information about metadata file
	Info struct {
		//Parent file metadata one related to
		GameRole

		Parent                   string
		ScrollableX, ScrollableY bool
	}

	Size Vec2

	Margins Vec2

	Spawns []*server_external.PositionInt

	Physics struct {
		G float64
	}

	Buffs struct {
		Speed Vec2
	}

	Scale Vec2

	Offset Vec2

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

func (m *MetadataModel) GetSizeMaxX() float64 {
	ms := m.GetMargins()
	s := m.GetSize()
	return (ms.X + s.X)
}

func (m *MetadataModel) GetSizeMaxY() float64 {
	ms := m.GetMargins()
	s := m.GetSize()
	return (ms.Y + s.Y)
}

func (m *MetadataModel) GetSizeMinX() float64 {
	return m.GetMargins().X
}

func (m *MetadataModel) GetSizeMinY() float64 {
	return m.GetMargins().Y
}

func (m *MetadataModel) GetSize() Vec2 {
	s := m.GetScale()
	return Vec2{X: m.Size.X * s.X, Y: m.Size.Y * s.Y}
}

func (m *MetadataModel) GetMargins() Vec2 {
	s := m.GetScale()
	r := Vec2{X: m.Margins.X * s.X, Y: m.Margins.Y * s.Y}

	if m.Type.Contains("scrollable") {
		o := m.GetOffset()

		if m.Type.Contains("scrollablex") {
			r.X += o.X
		}
		if m.Type.Contains("scrollabley") {
			r.Y += o.Y
		}
	}

	return r
}

func (m *MetadataModel) GetScale() Vec2 {
	s := screen.UseScreen()
	screenWidth, screenHeight := s.GetSize()
	screenLastWidth, screenLastHeight := s.GetLastSize()
	return Vec2{X: (m.Scale.X * screenWidth) / screenLastWidth, Y: (m.Scale.Y * screenHeight) / screenLastHeight}
}

func (m *MetadataModel) GetBuffSpeed() Vec2 {
	s := m.GetScale()
	return Vec2{X: m.Buffs.Speed.X * s.X, Y: m.Buffs.Speed.Y * s.Y}
}

func (m *MetadataModel) GetOffset() Vec2 {
	s := m.GetScale()
	return Vec2{X: m.Offset.X * s.X, Y: m.Offset.Y * s.Y}
}

type ModelCombination struct {
	Modified, Origin *MetadataModel
	*RuntimeDefined
}

type Metadata struct {
	sync.Mutex

	Collection map[string]*MetadataModel
}

func (m *Metadata) loadFile(fs embed.FS, path string) {
	if reg := regexp.MustCompile(`\.toml*$`); reg.MatchString(path) {
		var mm MetadataModel

		if _, err := toml.DecodeFS(fs, path, &mm); err != nil {
			logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
		}

		if len(mm.Info.Parent) != 0 {
			path := filepath.Join(utils.GetBasePath(path), mm.Info.Parent)
			img := UseSources().Images().GetImage(path)

			x, y := img.Size()

			if mm.Animation.FrameNum != 0 {
				x /= int(mm.Animation.FrameNum)
			}

			mm.Size.X = float64(x)
			mm.Size.Y = float64(y)
		}

		m.Lock()
		m.Collection[reg.Split(path, -1)[0]] = &mm
		m.Unlock()
	}
}

func (m *Metadata) Load(fs embed.FS, path string, s, wg *sync.WaitGroup) {
	NewParser(fs, path, m.loadFile).Parse()
	wg.Done()
	s.Done()
}

func (m *Metadata) GetMetadata(path string) *MetadataModel {
	path = filepath.Join("dist/metadata", path)

	metadata, ok := m.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("metadata with path '%s' not found", path))
	}

	return metadata
}

func NewMetadata() *Metadata {
	return &Metadata{Collection: make(map[string]*MetadataModel)}
}
