package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

//Checks if any gamepad is connected
func (g *GamepadPress) IsGamepadConnected()bool{
	return len(ebiten.GamepadIDs()) != 0
}

func NewGamepad() *Gamepad {
	return new(Gamepad)
}
