package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	GamepadUPButton                ebiten.GamepadButton = 0
	GamepadRIGHTButton             ebiten.GamepadButton = 1
	GamepadDOWNButton              ebiten.GamepadButton = 2
	GamepadLEFTButton              ebiten.GamepadButton = 3
	GamepadLEFTUPPERCLICKERButton  ebiten.GamepadButton = 4
	GamepadRIGHTUPPERCLICKERButton ebiten.GamepadButton = 5
	GamepadLEFTLOWERCLICKERButton  ebiten.GamepadButton = 6
	GamepadRIGHTLOWERCLICKERButton ebiten.GamepadButton = 7
)

type Gamepad struct{
	GamepadPress
}

type GamepadPress struct{}

func (g *GamepadPress) IsGamepadButtonPressedOnce(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button) {
			return true
		}
	}
	return false
}

func (g *GamepadPress) IsGamepadButtonPressed(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button) {
			return true
		}
	}
	return false
}

func (g *GamepadPress) AreGamepadButtonsCombinedOnce(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button1) && inpututil.IsGamepadButtonJustPressed(v, button2) {
			return true
		}
	}
	return false
}

func (g *GamepadPress) AreGamepadButtonsCombined(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button1) && ebiten.IsGamepadButtonPressed(v, button2) {
			return true
		}
	}
	return false
}

func NewGamepad() *Gamepad {
	return new(Gamepad)
}
