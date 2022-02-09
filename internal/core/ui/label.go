package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type LabelOpts struct {
	// Metadata     *sources.MetadataModel
	Position     types.Vec2
	FontDistance float64
	Text         string
	RowWidth     float64
	Font         *sources.Font
	Color        color.Color
}

type Label struct {
	Opts *LabelOpts
}

func (l *Label) Update() {}

func (l *Label) Render(sm *screen.ScreenManager) {
	l.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     l.Opts.Position,
		FontDistance: l.Opts.FontDistance,
		Color:        l.Opts.Color,
		RowWidth:     l.Opts.RowWidth,
	})
}

func NewLabel(opts *LabelOpts) Component {
	return &Label{opts}
}
