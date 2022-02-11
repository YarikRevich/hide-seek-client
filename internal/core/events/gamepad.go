package events

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GamepadEventManager struct{}

func (g *GamepadEventManager) IsGamepadConnected() bool {
	return len(ebiten.GamepadIDs()) != 0
}

func NewGamepadEventManager() *GamepadEventManager {
	return new(GamepadEventManager)
}

type GamepadPressEventManager struct{}

func (g *GamepadPressEventManager) IsGamepadButtonPressedOnce(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button) {
			return true
		}
	}
	return false
}

func (g *GamepadPressEventManager) IsGamepadButtonPressed(button ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button) {
			return true
		}
	}
	return false
}

func (g *GamepadPressEventManager) AreGamepadButtonsCombinedOnce(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if inpututil.IsGamepadButtonJustPressed(v, button1) && inpututil.IsGamepadButtonJustPressed(v, button2) {
			return true
		}
	}
	return false
}

func (g *GamepadPressEventManager) AreGamepadButtonsCombined(button1, button2 ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, button1) && ebiten.IsGamepadButtonPressed(v, button2) {
			return true
		}
	}
	return false
}

func (g *GamepadPressEventManager) AreGamepadButtonsCombinedInOrder(m, s ebiten.GamepadButton) bool {
	for _, v := range ebiten.GamepadIDs() {
		if ebiten.IsGamepadButtonPressed(v, m) == true && ebiten.IsGamepadButtonPressed(v, s) == true && inpututil.GamepadButtonPressDuration(v, m) > inpututil.GamepadButtonPressDuration(v, s) {
			return true
		}
	}
	return false
}

func (g *GamepadPressEventManager) IsAnyButtonPressed() bool {
	for _, v := range ebiten.GamepadIDs() {
		for _, q := range []ebiten.GamepadButton{
			keycodes.GamepadUPButton,
			keycodes.GamepadRIGHTButton,
			keycodes.GamepadDOWNButton,
			keycodes.GamepadLEFTButton,
			keycodes.GamepadLEFTUPPERCLICKERButton,
			keycodes.GamepadRIGHTUPPERCLICKERButton,
			keycodes.GamepadLEFTLOWERCLICKERButton,
			keycodes.GamepadRIGHTLOWERCLICKERButton} {
			if ebiten.IsGamepadButtonPressed(v, q) {
				return true
			}
		}
	}
	return false
}

func NewGamepadPressEventManager() *GamepadPressEventManager {
	return new(GamepadPressEventManager)
}

type GamepadScrollEventManager struct {
	IsMoved bool

	Offset, LastOffset types.Vec2

	Speed float64
}

func (gsem *GamepadScrollEventManager) UpdateGamepadScrollOffset() {
	if GamepadPress.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
		gsem.Offset.Y -= gsem.Speed
	} else if GamepadPress.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
		gsem.Offset.Y += gsem.Speed
	} else if GamepadPress.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
		gsem.Offset.X -= gsem.Speed
	} else if GamepadPress.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
		gsem.Offset.X += gsem.Speed
	}

	gsem.IsMoved = gsem.LastOffset.X != gsem.Offset.X && gsem.LastOffset.Y != gsem.Offset.Y
	gsem.LastOffset.X = gsem.Offset.X
	gsem.LastOffset.Y = gsem.Offset.Y
}

func NewGamepadScrollEventManager() *GamepadScrollEventManager {
	return &GamepadScrollEventManager{Speed: .5}
}
