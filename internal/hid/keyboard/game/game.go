package game

import (
	gamepadpress "github.com/YarikRevich/HideSeek-Client/internal/detectors/gamepad_press"
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {

	if gamepadpress.AreGamepadButtonsCombined(gamepadpress.GamepadUPButton, gamepadpress.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		camera.UseCamera().ZoomIn()
		return
	}

	if gamepadpress.AreGamepadButtonsCombined(gamepadpress.GamepadDOWNButton, gamepadpress.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		camera.UseCamera().ZoomOut()
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadUPButton) {
		if pc.UsePC().RawPos.Y > 0 {
			history.SetDirection(direction.UP)
			pc.UsePC().SetY(pc.UsePC().RawPos.Y - pc.UsePC().Buffs.SpeedY)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadDOWNButton) {
		history.SetDirection(direction.DOWN)
		pc.UsePC().SetY(pc.UsePC().RawPos.Y + pc.UsePC().Buffs.SpeedY)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadRIGHTButton) {
		history.SetDirection(direction.RIGHT)
		pc.UsePC().SetX(pc.UsePC().RawPos.X + pc.UsePC().Buffs.SpeedX)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadLEFTButton) {
		if pc.UsePC().RawPos.X > 0 {
			history.SetDirection(direction.LEFT)

			pc.UsePC().SetX(pc.UsePC().RawPos.X - pc.UsePC().Buffs.SpeedX)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) || gamepadpress.IsGamepadButtonPressed(gamepadpress.GamepadRIGHTUPPERCLICKERButton) {
		jump.CalculateJump(pc.UsePC())
	}

}
