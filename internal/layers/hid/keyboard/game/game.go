package game

import (
	"fmt"

	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/physics"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Exec() {

	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()

	c := world.UseWorld().GetCamera()
	// cPos := c.GetZoomedPos(&c.Base)
	p := world.UseWorld().GetPC()
	// pOffset := p.MetadataModel.GetOffset()
	pSpeed := p.MetadataModel.GetBuffSpeed()
	pScale := c.GetZoomedScale(&p.Base)
	wM := world.UseWorld().GetWorldMap()
	// wMScale := c.GetZoomedScale(&wM.Base)
	wMSize := wM.GetSize()

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

		fmt.Println(pSpeed)

		//Handles in-game minimap
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			switch statemachine.UseStateMachine().Minimap().GetState() {
			case statemachine.MINIMAP_OFF:
				statemachine.UseStateMachine().Minimap().SetState(statemachine.MINIMAP_ON)
			case statemachine.MINIMAP_ON:
				statemachine.UseStateMachine().Minimap().SetState(statemachine.MINIMAP_OFF)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if p.RawOffset.Y > -sHUD/4 {
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
			if p.RawOffset.Y < sAxis.Y*2 {
				p.SetRawY(p.RawPos.Y + pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y + pSpeed.Y)
				}

				if p.TranslationMovementYBlocked {
					c.SetRawY(c.RawPos.Y + pSpeed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			p.SetRawX(p.RawPos.X + pSpeed.X)
			if !p.TranslationMovementXBlocked {
				p.SetRawOffsetX(p.RawOffset.X + pSpeed.X)
			}

			if p.TranslationMovementXBlocked {
				c.SetRawX(c.RawPos.X + pSpeed.X)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if p.RawOffset.X > 0 {
				p.SetRawX(p.RawPos.X - pSpeed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - pSpeed.X)
				}

				if p.TranslationMovementXBlocked {
					c.SetRawX(c.RawPos.X - pSpeed.X)
				}
			}
		}

		// if p.TranslationMovementYBlocked {

		// 	if c.RawPos.Y <= -sHUD/2 && p.IsDirectionUP() {
		// 		c.SetRawY(0)
		// 		p.SetTranslationYMovementBlocked(false)
		// 	}

		// 	// fmt.Println(cPos, sAxis, wMSize, wMScale)
		fmt.Println(wMSize.Y+(sAxis.Y*2)+sHUD, c.RawPos.Y, "MAP RANGE")

		// 	fmt.Println(wMSize.Y, c.RawPos.Y)
		// 	if !c.IsOuttaRange(wMSize.Y, c.RawPos.Y) {
		// 		if c.RawPos.Y+sAxis.Y >= wMSize.Y && p.IsDirectionDOWN() {
		// 			// c.SetRawY(wMSize.Y - (sAxis.Y))
		// 			// fmt.Println(cPos, sAxis, wMSize, wMScale, "HERE")
		// 			p.SetTranslationYMovementBlocked(false)
		// 		}
		// 	}
		// }

		// if p.TranslationMovementXBlocked {
		// 	// wM.MetadataModel.GetSize()
		// 	// if cPos.X+sAxis.X*2 >= worldMap.Modified.Size.Width/worldMap.Modified.RuntimeDefined.ZoomedScale.X*1.81 &&
		// 	// 	p.IsDirectionRIGHT() {
		// 	// 	p.SetTranslationXMovementBlocked(false)
		// 	// }
		// 	if cPos.X <= 0 && p.IsDirectionLEFT() {
		// 		c.SetRawX(0)
		// 		p.SetTranslationXMovementBlocked(false)
		// 	}
		// }

		if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(p.RawOffset.X, pSpeed.X) && p.IsDirectionLEFT() {
			p.SetTranslationXMovementBlocked(true)
		}

		fmt.Println(p.RawOffset.X*pScale.X, pSpeed.X, s.IsHigherAxisXCrossed(p.RawOffset.X*pScale.X, pSpeed.X))
		if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(p.RawOffset.X, pSpeed.X) && p.IsDirectionRIGHT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(p.RawOffset.Y, pSpeed.Y) && p.IsDirectionUP() {
			p.SetTranslationYMovementBlocked(true)
		}

		// fmt.Println(s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y), pOffset.Y, pSpeed.Y)

		if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(p.RawOffset.Y, pSpeed.Y) && p.IsDirectionDOWN() {
			p.SetTranslationYMovementBlocked(true)
		}

		p.UpdateDirection()

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
