package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type LabelOpts struct {
	Metadata     *sources.MetadataModel
	Position     types.Vec2
	FontDistance float64
	Text         string
	RowWidth     float64
	Font         *sources.Font
	Color        color.Color
}

type Label struct {
	opts *LabelOpts
}

func (l *Label) Render(sm screen.ScreenManager) {
	l.opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     l.opts.Position,
		FontDistance: l.opts.FontDistance,
		Color:        l.opts.Color,
		RowWidth:     l.opts.RowWidth,
	})
}

func NewLabel(opts *LabelOpts) Component {
	return &Label{opts}
}
