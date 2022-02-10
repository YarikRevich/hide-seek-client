package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type ButtonOpts struct {
	// Metadata                  *sources.MetadataModel
	Tilemap                       *sources.Tilemap
	Position, Scale               types.Vec2
	AutoScaleForbidden            bool
	FontDistance, FontAdvance     float64
	Text                          string
	RowWidth                      float64
	Font                          *sources.Font
	Color                         color.Color
	OnMousePress, OnKeyboardPress func()
}

type Button struct {
	Opts *ButtonOpts
}

func (b *Button) Update() {
	if b.Opts.OnMousePress != nil {
		b.Opts.OnMousePress()
	}
	if b.Opts.OnKeyboardPress != nil {
		b.Opts.OnKeyboardPress()
	}
}

func (b *Button) Render(sm *screen.ScreenManager) {
	b.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		Position:           b.Opts.Position,
		Scale:              b.Opts.Scale,
		AutoScaleForbidden: b.Opts.AutoScaleForbidden,
	})
	b.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     b.Opts.Position,
		FontAdvance:  b.Opts.FontAdvance,
		FontDistance: b.Opts.FontDistance,
		Color:        b.Opts.Color,
		RowWidth:     b.Opts.RowWidth,
		Text:         b.Opts.Text,
	})
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{opts}
}
