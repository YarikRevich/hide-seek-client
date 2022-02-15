package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type BackgroundOpts struct {
	ID string

	Tilemap  *sources.Tilemap
	Position types.Vec2
}

type Background struct {
	Opts        *BackgroundOpts
	ContextOpts *ContextOpts
}

func (b *Background) SetContext(opts *ContextOpts) {
	b.ContextOpts = opts
}

func (b *Background) Update() {}

func (b *Background) Render(sm *screen.ScreenManager) {
	b.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:    types.Vec2{X: 0, Y: 0},
		AutoScaleForbidden: true,
	})
}

func (b *Background) GetID() string {
	return b.Opts.ID
}

func (b *Background) GetTilemap() *sources.Tilemap {
	return b.Opts.Tilemap
}

func (b *Background) GetPosition() types.Vec2 {
	return b.Opts.Position
}

func NewBackground(opts *BackgroundOpts) Component {
	return &Background{Opts: opts}
}
