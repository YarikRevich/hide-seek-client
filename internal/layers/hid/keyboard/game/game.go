package game

import (
	"math"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/physics"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()

	c := world.UseWorld().GetCamera()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		c.ZoomIn()
		return
	} else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		c.ZoomOut()
		return
	}

	if k.IsAnyKeyPressed() {
		worldMap := world.UseWorld().GetWorldMap()
		p := world.UseWorld().GetPC()

		pWidth, pHeight := p.ModelCombination.Modified.Size.Width, p.ModelCombination.Modified.Size.Height

		p.SaveLastPosition()
		pX, pY := p.GetScaledPosX(), p.GetScaledOffsetY()
		poX, poY := p.GetScaledOffsetX(), p.GetScaledOffsetY()
		cX, cY := c.GetScaledPosX(), c.GetScaledPosY()

		s := screen.UseScreen()

		hudOffsetHeight := s.GetHUDOffset()
		screenOffsetX := s.GetOffsetX()
		screenOffsetY := s.GetOffsetY()

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if pY > hudOffsetHeight {
				p.SetRawY(p.RawPos.Y - p.Modified.Buffs.Speed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y - p.Modified.Buffs.Speed.Y)
				}
			}

			if p.TranslationMovementYBlocked {
				c.SetRawY(c.RawPos.Y - p.Modified.Buffs.Speed.Y)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			if !p.TranslationMovementYBlocked {
				p.SetRawY(p.RawPos.Y + p.Modified.Buffs.Speed.Y)
				if pY < screenOffsetY*2-(pHeight/2) {
					p.SetRawOffsetY(p.RawOffset.Y + p.Modified.Buffs.Speed.Y)
				}
			}

			if p.TranslationMovementYBlocked {
				c.SetRawY(c.RawPos.Y + p.Modified.Buffs.Speed.Y)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			if !p.TranslationMovementXBlocked {
				p.SetRawX(p.RawPos.X + p.Modified.Buffs.Speed.X)
				if pX < screenOffsetX*2-(pWidth/2) {
					p.SetRawOffsetX(p.RawOffset.X + p.Modified.Buffs.Speed.X)
				}
			}

			if p.TranslationMovementXBlocked {
				c.SetRawX(c.RawPos.X + p.Modified.Buffs.Speed.X)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {

			if pX > 0 {
				p.SetRawX(p.RawPos.X - p.Modified.Buffs.Speed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - p.Modified.Buffs.Speed.X)
				}
			}

			if p.TranslationMovementXBlocked {
				c.SetRawX(c.RawPos.X - p.Modified.Buffs.Speed.X)
			}
		}

		if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(poX, pWidth) && p.IsDirectionLEFT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(poX, pWidth) && p.IsDirectionRIGHT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(poY, pHeight) && p.IsDirectionUP() {
			p.SetTranslationYMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(poY, pHeight) && p.IsDirectionDOWN() {
			p.SetTranslationYMovementBlocked(true)
		}

		p.UpdateDirection()

		if p.TranslationMovementYBlocked {
			if cY <= -hudOffsetHeight && p.IsDirectionUP() {
				p.SetTranslationYMovementBlocked(false)
			}

			if cY+screenOffsetY*2 >= worldMap.ModelCombination.Modified.Size.Height*worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.Y &&
				p.IsDirectionDOWN() {
				p.SetTranslationYMovementBlocked(false)
			}
		}

		if p.TranslationMovementXBlocked {
			// fmt.Println(cX, screenOffsetX, worldMap.ModelCombination.Modified.Size.Width*worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.X)
			if math.Ceil(cX+screenOffsetX*2) >= math.Ceil(worldMap.ModelCombination.Modified.Size.Width*worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.X) &&
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
