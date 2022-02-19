package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type LabelOpts struct {
	TextOpts

	ID string

	//Should contain ID to the ui object
	//you want this object to connect to
	StickedTo              string
	Tilemap                string
	SurfacePosition, Scale types.Vec2
}

type Label struct {
	Tilemap sources.Tilemap

	Opts        *LabelOpts
	ContextOpts *ContextOpts
}

func (l *Label) SetContext(opts *ContextOpts) {
	l.ContextOpts = opts
}

func (l *Label) Update(sm *screen.ScreenManager) {}

func (l *Label) Render(sm *screen.ScreenManager) {
	l.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:  l.Opts.Position,
		Scale:            l.Opts.Scale,
		CenterizedOffset: true,
	})

	l.Opts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Tilemap:         &l.Tilemap,
		SurfacePosition: l.Opts.SurfacePosition,
		TextPosition:    l.Opts.TextOpts.Position,
		Color:           l.Opts.Color,
		RowWidth:        l.Opts.RowWidth,
		Text:            l.Opts.Text,
	})
}

func (l *Label) GetID() string {
	return l.Opts.ID
}

func (l *Label) GetTilemap() *sources.Tilemap {
	return &l.Tilemap
}

func (l *Label) GetPosition() types.Vec2 {
	return l.Opts.Position
}

func NewLabel(opts *LabelOpts) Component {
	return &Label{
		Tilemap: sources.GetTileMap(opts.Tilemap),
		Opts:    opts}
}
