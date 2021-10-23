package gamepad_press

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Kube is 3
//Angle is 0
//Circle is 1
//Cross is 2
//Лівий нижній курок 6
//Лівий верхній курок 4
//Правий нижній курок 7
//Правий верхній курок 5
const (
	GamepadUPButton                ebiten.GamepadButton = 0
	GamepadRIGHTButton             ebiten.GamepadButton = 1
	GamepadDOWNButton              ebiten.GamepadButton = 2
	GamepadLEFTButton              ebiten.GamepadButton = 3
	GamepadLEFTUPPERCLICKERButton ebiten.GamepadButton = 4
	GamepadRIGHTUPPERCLICKERButton ebiten.GamepadButton = 5
	
)

func IsGamepadButtonPressedOnce(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button) {
			return true
		}
	}
	return false
}

func IsGamepadButtonPressed(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button) {
			return true
		}
	}
	return false
}

func AreGamepadButtonsCombinedOnce(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button1) && inpututil.IsGamepadButtonJustPressed(v, button2) {
			return true
		}
	}
	return false
}

func AreGamepadButtonsCombined(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button1) && ebiten.IsGamepadButtonPressed(v, button2) {
			return true
		}
	}
	return false
}
