package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type InputOpts struct {
	// Metadata     *sources.MetadataModel
	Position     types.Vec2
	FontDistance float64
	Text         string
	RowWidth     float64
	Font         *sources.Font
	Color        color.Color
}

type Input struct {
	Opts *InputOpts
}

func (in *Input) Render(sm *screen.ScreenManager) {
	in.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     in.Opts.Position,
		FontDistance: in.Opts.FontDistance,
		Color:        in.Opts.Color,
		RowWidth:     in.Opts.RowWidth,
	})
}

func NewInput(opts *InputOpts) Component {
	return &Input{opts}
}
