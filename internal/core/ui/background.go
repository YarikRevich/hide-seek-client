package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type BackgroundOpts struct {
	ID string

	Tilemap         string
	Position, Scale types.Vec2

	CameraAngle, CameraPitch, CameraZoom float64
	CameraPosition                       types.Vec3

	OrthigraphicProjection, PerspectiveProjection bool

	Camera *camera.Camera
}

type Background struct {
	Tilemap sources.Tilemap

	Opts        *BackgroundOpts
	ContextOpts *ContextOpts
}

func (b *Background) SetContext(opts *ContextOpts) {
	b.ContextOpts = opts
}

func (b *Background) Update(sm *screen.ScreenManager) {}

func (b *Background) Render(sm *screen.ScreenManager) {
	// x, y := ebiten.CursorPosition()
	b.Tilemap.Render(sm, sources.RenderTilemapOpts{
		SurfacePosition:        types.Vec2{X: 0, Y: 0},
		Scale:                  b.Opts.Scale,
		AutoScaleForbidden:     true,
		OrthigraphicProjection: b.Opts.OrthigraphicProjection,
		RenderTilemapOptsContext: sources.RenderTilemapOptsContext{
			Camera: b.Opts.Camera,
		},
		// CameraPosition: types.Vec3{X: float64(x), Z: float64(y)},
		// CameraZoom:     b.Opts.CameraZoom,
		// CameraPosition: b.Opts.CameraPosition,
		// CameraAngle:    b.Opts.CameraAngle,
		// CameraPitch:    b.Opts.CameraPitch,
	})
}

func (b *Background) GetID() string {
	return b.Opts.ID
}

func (b *Background) GetTilemap() *sources.Tilemap {
	return &b.Tilemap
}

func (b *Background) GetPosition() types.Vec2 {
	return b.Opts.Position
}

func NewBackground(opts *BackgroundOpts) Component {
	return &Background{
		Tilemap: sources.GetTileMap(opts.Tilemap),
		Opts:    opts}
}
