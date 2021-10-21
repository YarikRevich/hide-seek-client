package game

import "github.com/hajimehoshi/ebiten/v2"

//Kube is 3
//Angle is 0
//Circle is 1
//Cross is 2
//Лівий нижній курок 6
//Лівий верхній курок 4
//Правий нижній курок 7
//Правий верхній курок 5
const (
	gamepadUPButton ebiten.GamepadButton = 0
	gamepadRIGHTButton ebiten.GamepadButton = 1
	gamepadDOWNButton ebiten.GamepadButton = 2
	gamepadLEFTButton ebiten.GamepadButton = 3
	gamepadRIGHTUPPERCLICKERButton ebiten.GamepadButton = 5
)

func isGamepadButtonPressed(button ebiten.GamepadButton)bool{
	for _, v := range ebiten.GamepadIDs(){
		if ebiten.IsGamepadButtonPressed(v, button){
			return true
		}
	}
	return false
}