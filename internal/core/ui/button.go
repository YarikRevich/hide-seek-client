package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type ButtonOpts struct {
	// Metadata                  *sources.MetadataModel
	Tilemap                   *sources.Tilemap
	Position                  types.Vec2
	FontDistance, FontAdvance float64
	Text                      string
	RowWidth                  float64
	Font                      *sources.Font
	Color                     color.Color
	OnMousePress              func()
}

type Button struct {
	Opts *ButtonOpts
}

func (b *Button) Render(sm *screen.ScreenManager) {
	b.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		Position: b.Opts.Position,
	})
	b.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Position:     b.Opts.Position,
		FontAdvance:  b.Opts.FontAdvance,
		FontDistance: b.Opts.FontDistance,
		Color:        b.Opts.Color,
		RowWidth:     b.Opts.RowWidth,
	})
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{opts}
}
