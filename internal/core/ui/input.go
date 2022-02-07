package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type InputOpts struct {
	Metadata     *sources.MetadataModel
	Position     types.Vec2
	FontDistance float64
	Text         string
	RowWidth     float64
	Font         *sources.Font
	Color        color.Color
}

type Input struct {
	opts *InputOpts
}

func (in *Input) Render(sm screen.ScreenManager) {
	in.opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     in.opts.Position,
		FontDistance: in.opts.FontDistance,
		Color:        in.opts.Color,
		RowWidth:     in.opts.RowWidth,
	})
}

func NewInput(opts *InputOpts) Component {
	return &Input{opts}
}
