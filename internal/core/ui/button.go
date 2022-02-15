package ui

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
)

type ButtonOpts struct {
	TextOpts TextOpts

	ID string

	//Should contain ID to the ui object
	//you want this object to connect to
	StickedTo              string
	Tilemap                *sources.Tilemap
	SurfacePosition, Scale types.Vec2

	OnMousePress, OnKeyboardPress func()
}

type Button struct {
	Opts        *ButtonOpts
	ContextOpts *ContextOpts
}

func (b *Button) SetContext(opts *ContextOpts) {
	b.ContextOpts = opts
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
	fmt.Println(b.ContextOpts.Components)

	renderTilemapsOpts := sources.RenderTilemapOpts{
		SurfacePosition:  b.Opts.SurfacePosition,
		Scale:            b.Opts.Scale,
		CenterizedOffset: true,
	}

	if stickedTo, ok := b.ContextOpts.Components[b.Opts.StickedTo]; ok {
		renderTilemapsOpts.StickedTo = stickedTo.GetTilemap()
		renderTilemapsOpts.StickedToPosition = stickedTo.GetPosition()
	}

	b.Opts.Tilemap.Render(sm, renderTilemapsOpts)

	b.Opts.TextOpts.Font.Render(sm, sources.RenderTextCharachterOpts{
		Align:           b.Opts.TextOpts.Align,
		Tilemap:         b.Opts.Tilemap,
		SurfacePosition: b.Opts.SurfacePosition,
		SurfaceScale:    b.Opts.Scale,
		TextPosition:    b.Opts.TextOpts.Position,
		Color:           b.Opts.TextOpts.Color,
		RowWidth:        b.Opts.TextOpts.RowWidth,
		Text:            b.Opts.TextOpts.Text,
	})
}

func (b *Button) GetID() string {
	return b.Opts.ID
}

func (b *Button) GetTilemap() *sources.Tilemap {
	return b.Opts.Tilemap
}

func (b *Button) GetPosition() types.Vec2 {
	return b.Opts.SurfacePosition
}

func NewButton(opts *ButtonOpts) Component {
	return &Button{Opts: opts}
}
