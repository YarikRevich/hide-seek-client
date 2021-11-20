package events

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	MousePress
	MouseWheel
}

type MousePress struct{}

func (p *MousePress) IsMousePressLeftOnce(m sources.Model) bool {
	currX, currY := ebiten.CursorPosition()
	mx, my := m.Margins.LeftMargin * m.Scale.X, m.Margins.TopMargin * m.Scale.X
	
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		(currX >= int(mx) && currX <= int((m.Size.Width)+(mx))) &&
		(currY >= int(my) && currY <= int((m.Size.Height)+(my)))
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

type MouseWheel struct {
	IsMoved bool

	LastMouseWheelX, LastMouseWheelY float64
	OffsetX, OffsetY         float64
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
		p.OffsetX += sx; p.OffsetY += sy
	}

	p.IsMoved = p.LastMouseWheelX != p.OffsetX && p.LastMouseWheelY != p.OffsetY
	p.LastMouseWheelX = p.OffsetX; p.LastMouseWheelY = p.OffsetY
}

func NewMouse() *Mouse {
	m := new(Mouse)
	m.moveCoefficient = .5
	return m
}
