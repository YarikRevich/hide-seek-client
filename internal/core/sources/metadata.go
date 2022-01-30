package sources

import (
	"embed"
	"fmt"
	"image"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
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
	ZoomedScale types.Vec2
}

type Transition struct {
	StartScale, EndScale types.Vec2
}

//All the positioning properties should be in range (0; 100)
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

	Size types.Vec2

	Margins types.Vec2

	Spawns []*server_external.PositionInt

	Physics struct {
		G float64
	}

	Buffs struct {
		Speed float64
	}

	Scale types.Vec2

	Offset types.Vec2

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

func (m *MetadataModel) GetRect() image.Rectangle {
	ms := m.GetMargins()
	s := m.GetSize()
	ma := m.GetMargins()
	return image.Rect(int(ma.X), int(ma.Y), int(ms.X+s.X), int(ms.Y+s.Y))
}

func (m *MetadataModel) GetSize() types.Vec2 {
	return m.Size
}

func (m *MetadataModel) GetMargins() types.Vec2 {
	ss := m.GetSize()
	sc := m.GetScale()
	s := screen.UseScreen()
	size := s.GetSize()
	lastSize := s.GetLastSize()
	r := types.Vec2{X: (((m.Margins.X * size.X) / 100) / (size.X / lastSize.X)) - (ss.X * sc.X / 2), Y: (((m.Margins.Y * size.Y) / 100) / (size.Y / lastSize.Y)) - (ss.Y * sc.Y / 2)}

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

func (m *MetadataModel) GetScale() types.Vec2 {
	s := screen.UseScreen()
	size := s.GetSize()
	lastSize := s.GetLastSize()
	return types.Vec2{X: (((m.Scale.X * size.X) / 100) / (size.X / lastSize.X)), Y: (((m.Scale.Y * size.Y) / 100) / (size.Y / lastSize.Y))}
}

func (m *MetadataModel) GetBuffSpeed() types.Vec2 {
	s := screen.UseScreen()
	size := s.GetSize()
	lastSize := s.GetLastSize()
	y := m.Buffs.Speed * (size.Y / lastSize.Y)
	x := (m.Buffs.Speed * (size.X / lastSize.X)) - (y / 2)
	avg := (y + x) / 2
	return types.Vec2{X: avg, Y: avg}
}

func (m *MetadataModel) GetOffset() types.Vec2 {
	s := screen.UseScreen()
	size := s.GetSize()
	lastSize := s.GetLastSize()
	return types.Vec2{X: m.Offset.X * (size.X / lastSize.X), Y: m.Offset.Y * (size.Y / lastSize.Y)}
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
			path := filepath.Join(utils.GetBasePath(strings.Split(path, "dist/metadata/")[1]), mm.Info.Parent)
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
