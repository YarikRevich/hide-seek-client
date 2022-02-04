package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"golang.org/x/image/font"
)

type ButtonOpts struct {
	Metadata                  *sources.MetadataModel
	Position                  types.Vec2
	FontDistance, FontAdvance float64
	Text                      string
	RowWidth                  float64
	Font                      font.Face
	Color                     color.Color
	OnMousePress              func()
}

type Button struct {
	opts *ButtonOpts
}

func (b *Button) Render(s screen.Screen) {
	for i, c := range b.opts.Text {
		s.RenderTextCharachter(i, c, screen.RenderTextCharachterOpts{
			Position:     b.opts.Position,
			FontAdvance:  b.opts.FontAdvance,
			FontDistance: b.opts.FontDistance,
			Font:         b.opts.Font,
			Color:        b.opts.Color,
			RowWidth:     b.opts.RowWidth,
		})
	}
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{opts}
}
