package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	if areGamepadButtonsCombined(gamepadUPButton, gamepadLEFTUPPERCLICKERButton){
		camera.UseCamera().ZoomIn()
		return
	}

	if areGamepadButtonsCombined(gamepadDOWNButton, gamepadLEFTUPPERCLICKERButton){
		camera.UseCamera().ZoomOut()
		return
	}


	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || isGamepadButtonPressed(gamepadUPButton) {
		history.SetDirection(direction.UP)
		pc.UsePC().SetY(pc.UsePC().RawPos.Y - pc.UsePC().Buffs.SpeedY)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || isGamepadButtonPressed(gamepadDOWNButton) {
		history.SetDirection(direction.DOWN)
		pc.UsePC().SetY(pc.UsePC().RawPos.Y + pc.UsePC().Buffs.SpeedY)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || isGamepadButtonPressed(gamepadRIGHTButton) {
		history.SetDirection(direction.RIGHT)
		pc.UsePC().SetX(pc.UsePC().RawPos.X + pc.UsePC().Buffs.SpeedX)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || isGamepadButtonPressed(gamepadLEFTButton) {
		history.SetDirection(direction.LEFT)
		pc.UsePC().SetX(pc.UsePC().RawPos.X - pc.UsePC().Buffs.SpeedX)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) || isGamepadButtonPressed(gamepadRIGHTUPPERCLICKERButton) {
		jump.CalculateJump(pc.UsePC())
	}

	
}
