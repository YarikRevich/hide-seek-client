package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/physics"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()
	
	c := objects.UseObjects().Camera()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		c.ZoomIn()
		return
	} else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		c.ZoomOut()
		return
	}

	if k.IsAnyKeyPressed() {
		w := objects.UseObjects().World()
		p := objects.UseObjects().PC()

		p.SaveLastPosition()

		pX, pY := p.GetScaledPosX(), p.GetScaledOffsetY()
		cX, cY := c.GetScaledPosX(), c.GetScaledPosY()
		// pZoomedSpeedX, pZoomedSpeedY := pmm.Buffs.Speed.X*mapScaleX, pmm.Buffs.Speed.Y*mapScaleY
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {

			if pY > 0 {
				p.SetRawY(p.RawPos.Y - p.Modified.Buffs.Speed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y - p.Modified.Buffs.Speed.Y)
				}
			}

			if cY > 0 {
				c.SetRawY(c.RawPos.Y - p.Modified.Buffs.Speed.Y)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			// if zy+pmm.Buffs.Speed.Y+pmo.Size.Height < wm.Size.Height*wsy {
			// 	p.SetRawY(p.RawPos.Y + pmm.Buffs.Speed.Y)
			// 	if !p.TranslationMovementYBlocked {
			// 		p.SetRawPosForCameraY(p.RawPosForCamera.Y + pmm.Buffs.Speed.Y)
			// 	}else{
			// 		c.SetRawY(c.RawPos.Y + pmm.Buffs.Speed.Y)
			// 	}
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			//+pmm.Buffs.Speed.X+pmo.Size.Width
			// fmt.Println(p.GetX(), c.GetX(), wm.Size.Width*wsx)
			// if p.GetRawX() < wm.Size.Width*wsx {

			// fmt.Println(pZoomedSpeedX)
			p.SetRawX(p.RawPos.X + p.Modified.Buffs.Speed.X)
			if !p.TranslationMovementXBlocked {
				p.SetRawOffsetX(p.RawOffset.X + p.Modified.Buffs.Speed.X)
			}
			// }

			// fmt.Println(cZoomedX, wm.Size.Width*wsx, cZoomedX < wm.Size.Width*wsx)
			// if c.GetRawX() < wm.Size.Width*wsx {
			// fmt.Println(cZoomedX, pZoomedOffsetX*2, pmo.Size.Width, pmm.Buffs.Speed.X < wm.Size.Width*mapScaleX, "PUK")
			if p.TranslationMovementXBlocked {
				c.SetRawX(c.RawPos.X + p.Modified.Buffs.Speed.X)
			}
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {

			if pX > 0 {
				p.SetRawX(p.RawPos.X - p.Modified.Buffs.Speed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - p.Modified.Buffs.Speed.X)
				}
			}

			if cX > 0 {
				if p.TranslationMovementXBlocked {
					c.SetRawX(c.RawPos.X - p.Modified.Buffs.Speed.X)
				}
			}
		}
		s := screen.UseScreen()
		if !p.TranslationMovementXBlocked && s.IsAxisXCrossedByPC() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsAxisYCrossedByPC() {
			p.SetTranslationYMovementBlocked(true)
		}

		p.UpdateDirection()

		poX := p.GetScaledPosX()

		if p.TranslationMovementYBlocked {
			if cY <= 0 && p.IsDirectionUP() {
				p.SetTranslationYMovementBlocked(false)
			}

			if cY >= w.ModelCombination.Modified.Size.Height*w.ModelCombination.Modified.ZoomedScale.Y &&
				p.IsDirectionDOWN() {
				p.SetTranslationYMovementBlocked(false)
			}
		}

		if p.TranslationMovementXBlocked {
			if cX+poX*2 >= w.ModelCombination.Modified.Size.Width*w.ModelCombination.Modified.ZoomedScale.X &&
				p.IsDirectionRIGHT() {
				p.SetTranslationXMovementBlocked(false)
			}
			if cX <= 0 && p.IsDirectionLEFT() {
				p.SetTranslationXMovementBlocked(false)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
