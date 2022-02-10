package events

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	MousePress
	MouseWheel
}

type MousePress struct{}

func (p *MousePress) IsMousePressLeftOnce(sm screen.ScreenManager, tm sources.Tilemap, b ui.ButtonOpts) bool {
	currX, currY := ebiten.CursorPosition()
	screenScale := sm.GetScale()

	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(b.Position.X) && currX <= int((tm.MapSize.X*screenScale.X)+(b.Position.X))) &&
		(currY >= int(b.Position.Y) && currY <= int((tm.MapSize.Y*screenScale.Y)+(b.Position.Y)))
}

//It checks collision with a static object, which won't change its size
//after window resizing
func (p *MousePress) IsMousePressLeftOnceStatic(tm sources.Tilemap, b ui.ButtonOpts) bool {
	currX, currY := ebiten.CursorPosition()

	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(b.Position.X) && currX <= int((tm.MapSize.X)+(b.Position.X))) &&
		(currY >= int(b.Position.Y) && currY <= int((tm.MapSize.Y)+(b.Position.Y)))
}

func (p *Mouse) IsAnyMouseButtonsPressed() bool {
	for _, v := range []ebiten.MouseButton{
		ebiten.MouseButtonLeft, ebiten.MouseButtonMiddle, ebiten.MouseButtonRight} {
		if inpututil.IsMouseButtonJustPressed(v) {
			return true
		}
	}
	return false
}

func (p *Mouse) IsAnyMovementButtonPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyS) ||
		ebiten.IsKeyPressed(ebiten.KeyA) ||
		ebiten.IsKeyPressed(ebiten.KeyD) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowRight)
}

type MouseWheel struct {
	IsMoved bool

	LastMouseWheelX, LastMouseWheelY float64
	OffsetX, OffsetY                 float64
	moveCoefficient                  float64
}

//Saves mouse wheel offsets using ebiten API
//or uses offsets gotten from gamepad
func (p *MouseWheel) UpdateMouseWheelOffsets() {
	e := UseEvents().Gamepad()

	if e.IsGamepadConnected() {
		if e.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			p.OffsetY -= p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			p.OffsetY += p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			p.OffsetX -= p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			p.OffsetX += p.moveCoefficient
		}
	} else {
		sx, sy := ebiten.Wheel()
		p.OffsetX += sx
		p.OffsetY += sy
	}

	p.IsMoved = p.LastMouseWheelX != p.OffsetX && p.LastMouseWheelY != p.OffsetY
	p.LastMouseWheelX = p.OffsetX
	p.LastMouseWheelY = p.OffsetY
}

func NewMouse() *Mouse {
	m := new(Mouse)
	m.moveCoefficient = .5
	return m
}
