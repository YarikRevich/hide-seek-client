package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type ButtonOpts struct {
	Tilemap                       *sources.Tilemap
	SurfacePosition, Scale        types.Vec2
	AutoScaleForbidden            bool
	FontDistance, FontAdvance     float64
	Text                          string
	TextPosition                  types.Vec2
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
		if events.MousePress.IsAnyMouseButtonsPressed() {
			// if events.MousePress.IsMousePressLeftOnce(b.Opts.)
			//TODO: check if button is pressed, then do action
			b.Opts.OnMousePress()
		}

	}
	if b.Opts.OnKeyboardPress != nil {
		b.Opts.OnKeyboardPress()
	}
}

func (b *Button) Render(sm *screen.ScreenManager) {
	b.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:    b.Opts.SurfacePosition,
		Scale:              b.Opts.Scale,
		AutoScaleForbidden: b.Opts.AutoScaleForbidden,
	})
	b.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		SurfacePosition: b.Opts.SurfacePosition,
		FontAdvance:     b.Opts.FontAdvance,
		FontDistance:    b.Opts.FontDistance,
		TextPosition:    b.Opts.TextPosition,
		Color:           b.Opts.Color,
		RowWidth:        b.Opts.RowWidth,
		Text:            b.Opts.Text,
	})
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{opts}
}
