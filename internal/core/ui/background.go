package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type BackgroundOpts struct {
	Tilemap            *sources.Tilemap
	AutoScaleForbidden bool
}

type Background struct {
	Opts *BackgroundOpts
}

func (b *Background) Update() {}

func (b *Background) Render(sm *screen.ScreenManager) {
	b.Opts.Tilemap.Render(sm, sources.RenderTilemapOpts{
		Position:           types.Vec2{X: 0, Y: 0},
		AutoScaleForbidden: b.Opts.AutoScaleForbidden,
	})
}

func NewBackground(opts *BackgroundOpts) Component {
	return &Background{opts}
}
