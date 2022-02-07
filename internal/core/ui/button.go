package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type ButtonOpts struct {
	Metadata                  *sources.MetadataModel
	Position                  types.Vec2
	FontDistance, FontAdvance float64
	Text                      string
	RowWidth                  float64
	Font                      *sources.Font
	Color                     color.Color
	OnMousePress              func()
}

type Button struct {
	opts *ButtonOpts
}

func (b *Button) Render(sm screen.ScreenManager) {
	b.opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     b.opts.Position,
		FontAdvance:  b.opts.FontAdvance,
		FontDistance: b.opts.FontDistance,
		Color:        b.opts.Color,
		RowWidth:     b.opts.RowWidth,
	})
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{opts}
}
