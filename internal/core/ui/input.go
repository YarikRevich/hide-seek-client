package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type InputOpts struct {
	TextOpts

	ID string

	//Should contain ID to the ui object
	//you want this object to connect to
	StickedTo                     string
	Tilemap                       *sources.Tilemap
	SurfacePosition, Scale        types.Vec2
	OnMousePress, OnKeyboardPress func()
}

type Input struct {
	Opts        *InputOpts
	ContextOpts *ContextOpts
}

func (i *Input) SetContext(opts *ContextOpts) {
	i.ContextOpts = opts
}

func (i *Input) Update(sm *screen.ScreenManager) {
	i.Opts.OnMousePress()
	i.Opts.OnKeyboardPress()
}

func (in *Input) Render(sm *screen.ScreenManager) {
	in.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:  in.Opts.Position,
		Scale:            in.Opts.Scale,
		CenterizedOffset: true,
	})

	in.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Tilemap:         in.Opts.Tilemap,
		SurfacePosition: in.Opts.SurfacePosition,
		TextPosition:    in.Opts.TextOpts.Position,
		Color:           in.Opts.Color,
		RowWidth:        in.Opts.RowWidth,
		Text:            in.Opts.Text,
	})
}

func (in *Input) GetID() string {
	return in.Opts.ID
}

func (in *Input) GetTilemap() *sources.Tilemap {
	return in.Opts.Tilemap
}

func (in *Input) GetPosition() types.Vec2 {
	return in.Opts.Position
}

func NewInput(opts *InputOpts) Component {
	return &Input{Opts: opts}
}
