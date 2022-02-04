package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"golang.org/x/image/font"
)

type InputOpts struct {
	Metadata     *sources.MetadataModel
	Position     types.Vec2
	FontDistance float64
	Text         string
	RowWidth     float64
	Font         font.Face
	Color        color.Color
}

type Input struct {
	opts *InputOpts
}

func (in *Input) Render(s screen.Screen) {
	for i, c := range in.opts.Text {
		s.RenderTextCharachter(i, c, screen.RenderTextCharachterOpts{
			Position:     in.opts.Position,
			FontDistance: in.opts.FontDistance,
			Font:         in.opts.Font,
			Color:        in.opts.Color,
			RowWidth:     in.opts.RowWidth,
		})
	}
}

func NewInput(opts *InputOpts) Component {
	return &Input{opts}
}
