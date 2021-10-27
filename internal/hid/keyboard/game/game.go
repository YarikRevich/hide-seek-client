package game

import (
	// gamepadpress "github.com/YarikRevich/HideSeek-Client/internal/detectors/gamepad_press"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	// "github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	g := events.UseEvents().Gamepad()
	p := objects.UseObjects().PC()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		camera.UseCamera().ZoomIn()
		return
	}

	if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		camera.UseCamera().ZoomOut()
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
		if p.RawPos.Y > 0 {
			p.SetY(p.RawPos.Y - p.Buffs.SpeedY)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
	
		p.SetY(p.RawPos.Y + p.Buffs.SpeedY)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
	
		p.SetX(p.RawPos.X + p.Buffs.SpeedX)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
		if p.RawPos.X > 0 {
	
			p.SetX(p.RawPos.X - p.Buffs.SpeedX)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
		jump.CalculateJump(&p.Object)
	}

}
