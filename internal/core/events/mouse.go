package events

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	MousePress
	MouseWheel
}

type MousePress struct{}

func (p *MousePress) IsMousePressLeftOnce(m models.Metadata) bool {
	currX, currY := ebiten.CursorPosition()
	mx, my := m.FastenMarginsWithCoefficients()

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
	MouseWheelX, MouseWheelY         float64
	moveCoefficient                  float64
}

//Saves mouse wheel offsets using ebiten API
//or uses offsets gotten from gamepad
func (p *MouseWheel) UpdateMouseWheelOffsets() {
	e := UseEvents().Gamepad()

	if e.IsGamepadConnected() {
		if e.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			p.MouseWheelY -= p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			p.MouseWheelY += p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			p.MouseWheelX -= p.moveCoefficient
		} else if e.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			p.MouseWheelX += p.moveCoefficient
		}
	} else if sx, sy := ebiten.Wheel(); sx != 0 || sy != 0 {
		p.MouseWheelX += sx; p.MouseWheelY += sy
	}

	p.IsMoved = p.LastMouseWheelX != p.MouseWheelX && p.LastMouseWheelY != p.MouseWheelY
	p.LastMouseWheelX = p.MouseWheelX; p.LastMouseWheelY = p.MouseWheelY
}

func NewMouse() *Mouse {
	m := new(Mouse)
	m.moveCoefficient = 2
	return m
}
