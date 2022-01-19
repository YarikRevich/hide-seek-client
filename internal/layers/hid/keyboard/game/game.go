package game

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/physics"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	// wM := world.UseWorld().GetWorldMap()
	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()

	c := world.UseWorld().GetCamera()
	// cPos := c.GetZoomedPos(&c.Base)
	p := world.UseWorld().GetPC()
	pOffset := c.GetZoomedOffset(&p.Base)
	pSpeed := p.MetadataModel.GetBuffSpeed()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		c.ZoomIn(&p.Base)
		p.UpdateLastActivity()
		return
	} else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		c.ZoomOut(&p.Base)
		p.UpdateLastActivity()
		return
	}

	if k.IsAnyKeyPressed() || g.IsAnyButtonPressed() {
		p.UpdateLastActivity()
		p.UpdateLastPosition()

		s := screen.UseScreen()
		sAxis := s.GetAxis()
		sHUD := s.GetHUDOffset()

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if pOffset.Y > -sHUD {
				p.SetRawY(p.RawPos.Y - pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y - pSpeed.Y)
				}
				if p.TranslationMovementYBlocked {
					c.SetRawY(c.RawPos.Y - pSpeed.Y)
				}
			}

		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			if pOffset.Y < sAxis.Y*2 {
				p.SetRawY(p.RawPos.Y + pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					fmt.Println("HERe", p.RawOffset.Y)
					p.SetRawOffsetY(p.RawOffset.Y + pSpeed.Y)
				}

				if p.TranslationMovementYBlocked {
					c.SetRawY(c.RawPos.Y + pSpeed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			if pOffset.X < sAxis.X*2 {
				p.SetRawX(p.RawPos.X + pSpeed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X + pSpeed.X)
				}

				if p.TranslationMovementXBlocked {
					c.SetRawX(c.RawPos.X + pSpeed.X)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if pOffset.X > 0 {
				p.SetRawX(p.RawPos.X - pSpeed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - pSpeed.X)
				}

				if p.TranslationMovementXBlocked {
					c.SetRawX(c.RawPos.X - pSpeed.X)
				}
			}
		}

		if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionLEFT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionRIGHT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionUP() {
			p.SetTranslationYMovementBlocked(true)
		}

		fmt.Println(s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y), pOffset.Y, pSpeed.Y)

		if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionDOWN() {
			p.SetTranslationYMovementBlocked(true)
		}

		p.UpdateDirection()

		if p.TranslationMovementYBlocked {
			// if cPos.Y <= -sHUD && p.IsDirectionUP() {
			// 	p.SetTranslationYMovementBlocked(false)
			// }

			// if cPos.Y+sAxis.Y*2 >= worldMap.Modified.Size.Height/1.81 &&
			// 	p.IsDirectionDOWN() {
			// 	p.SetTranslationYMovementBlocked(false)
			// }
		}

		if p.TranslationMovementXBlocked {
			// wM.MetadataModel.GetSize()
			// if cPos.X+sAxis.X*2 >= worldMap.Modified.Size.Width/worldMap.Modified.RuntimeDefined.ZoomedScale.X*1.81 &&
			// 	p.IsDirectionRIGHT() {
			// 	p.SetTranslationXMovementBlocked(false)
			// }
			// if cPos.X <= 0 && p.IsDirectionLEFT() {
			// 	p.SetTranslationXMovementBlocked(false)
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
