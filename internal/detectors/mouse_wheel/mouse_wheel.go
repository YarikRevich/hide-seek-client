package mousewheel

import (
	gamepadpress "github.com/YarikRevich/HideSeek-Client/internal/detectors/gamepad_press"
	"github.com/hajimehoshi/ebiten/v2"
)

var MouseWheelX, MouseWheelY float64

const moveCoefficient = 2

//Saves mouse wheel offsets using ebiten API
//or uses offsets gotten from gamepad
func SaveMouseWheelOffsets() {
	if gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadUPButton) {
		MouseWheelY -= moveCoefficient
	} else if gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadDOWNButton) {
		MouseWheelY += moveCoefficient
	} else if gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadLEFTButton) {
		MouseWheelX -= moveCoefficient
	} else if gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadRIGHTButton) {
		MouseWheelX += moveCoefficient
	} else {
		sx, sy := ebiten.Wheel()
		MouseWheelX += sx
		MouseWheelY += sy
	}
}
