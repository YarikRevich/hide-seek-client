package events

import (
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
	MouseWheelX, MouseWheelY float64
	moveCoefficient          float64
}

//Saves mouse wheel offsets using ebiten API
//or uses offsets gotten from gamepad
func (p *MouseWheel) UpdateMouseWheelOffsets() {
	e := UseEvents().Gamepad()

	if e.IsGamepadButtonPressed(GamepadUPButton) {
		p.MouseWheelY -= p.moveCoefficient
	} else if e.IsGamepadButtonPressed(GamepadDOWNButton) {
		p.MouseWheelY += p.moveCoefficient
	} else if e.IsGamepadButtonPressed(GamepadLEFTButton) {
		p.MouseWheelX -= p.moveCoefficient
	} else if e.IsGamepadButtonPressed(GamepadRIGHTButton) {
		p.MouseWheelX += p.moveCoefficient
	} else {
		sx, sy := ebiten.Wheel()
		p.MouseWheelX += sx
		p.MouseWheelY += sy
	}
}

func NewMouse() *Mouse {
	m := new(Mouse)
	m.moveCoefficient = 2
	return m
}
