package events

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// type Mouse struct {
// 	MouseWheel
// }

type MousePressEventManager struct{}

type IsMousePress struct {
	Position, MapSize, MapScale types.Vec2
}

func (mpem *MousePressEventManager) IsMousePressLeftOnce(sm *screen.ScreenManager, opts IsMousePress) bool {
	currX, currY := ebiten.CursorPosition()
	screenScale := sm.GetScale()
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(opts.Position.X-(opts.MapSize.X/2*opts.MapScale.X)) && currX <= int((opts.MapSize.X/2*opts.MapScale.X/screenScale.X)+(opts.Position.X))) &&
		(currY >= int(opts.Position.Y-(opts.MapSize.Y/2*opts.MapScale.Y)) && currY <= int((opts.MapSize.Y/2*opts.MapScale.Y/screenScale.Y)+(opts.Position.Y)))
}

//It checks collision with a static object, which won't change its size
//after window resizing
func (mpem *MousePressEventManager) IsMousePressLeftOnceStatic(tm sources.Tilemap, b IsMousePress) bool {
	currX, currY := ebiten.CursorPosition()

	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(b.Position.X) && currX <= int((tm.MapSize.X)+(b.Position.X))) &&
		(currY >= int(b.Position.Y) && currY <= int((tm.MapSize.Y)+(b.Position.Y)))
}

func (mpem *MousePressEventManager) IsAnyMouseButtonsPressed() bool {
	for _, v := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft, ebiten.MouseButtonMiddle, ebiten.MouseButtonRight} {
		if inpututil.IsMouseButtonJustPressed(v) {
			Activity.SetLastActivity()
			return true
		}
	}
	return false
}

func (mpem *MousePressEventManager) IsAnyMovementButtonPressed() bool {
	r := ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyS) ||
		ebiten.IsKeyPressed(ebiten.KeyA) ||
		ebiten.IsKeyPressed(ebiten.KeyD) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	if r {
		Activity.SetLastActivity()
	}
	return r
}

func NewMousePressEventManager() *MousePressEventManager {
	return new(MousePressEventManager)
}

type MouseScrollEventManager struct {
	IsMoved bool

	Offset, LastOffset types.Vec2
	Speed              float64
}

//Saves mouse wheel offsets using ebiten API
//or uses offsets gotten from gamepad
func (msem *MouseScrollEventManager) UpdateMouseScrollOffsets() {
	sx, sy := ebiten.Wheel()
	msem.Offset.X += (sx * msem.Speed)
	msem.Offset.Y += (sy * msem.Speed)

	msem.IsMoved = msem.LastOffset.X != msem.Offset.X && msem.LastOffset.Y != msem.Offset.Y
	msem.LastOffset.X = msem.Offset.X
	msem.LastOffset.Y = msem.Offset.Y
}

func NewMouseScrollEventManager() *MouseScrollEventManager {
	return &MouseScrollEventManager{Speed: 1}
}
