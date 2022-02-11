package ui

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type InputOpts struct {
	Tilemap                       *sources.Tilemap
	Position, Scale               types.Vec2
	AutoScaleForbidden            bool
	FontDistance                  float64
	Text                          string
	RowWidth                      float64
	Font                          *sources.Font
	Color                         color.Color
	OnMousePress, OnKeyboardPress func()
}

type Input struct {
	Opts *InputOpts
}

func (i *Input) Update() {
	i.Opts.OnMousePress()
	i.Opts.OnKeyboardPress()
}

func (in *Input) Render(sm *screen.ScreenManager) {
	in.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:    in.Opts.Position,
		Scale:              in.Opts.Scale,
		AutoScaleForbidden: in.Opts.AutoScaleForbidden,
	})
	in.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		SurfacePosition: in.Opts.Position,
		FontDistance:    in.Opts.FontDistance,
		Color:           in.Opts.Color,
		RowWidth:        in.Opts.RowWidth,
	})
}

func NewInput(opts *InputOpts) Component {
	return &Input{opts}
}
