package game

import (
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
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
	p := world.UseWorld().GetPC()
	c := camera.UseCamera()

	pSpeed := p.MetadataModel.GetBuffSpeed()
	wM := world.UseWorld().GetWorldMap()
	wMSize := wM.GetSize()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		c.Zoom(1.1)
		p.UpdateLastActivity()
		return
	} else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		c.Zoom(0.9)
		p.UpdateLastActivity()
		return
	}

	if k.IsAnyKeyPressed() || g.IsAnyButtonPressed() {
		p.UpdateLastActivity()
		p.UpdateLastPosition()

		s := screen.UseScreen()
		sAxis := s.GetAxis()
		sHUD := s.GetHUDOffset()

		//Handles in-game minimap
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			switch statemachine.UseStateMachine().Minimap().GetState() {
			case statemachine.MINIMAP_OFF:
				statemachine.UseStateMachine().Minimap().SetState(statemachine.MINIMAP_ON)
			case statemachine.MINIMAP_ON:
				statemachine.UseStateMachine().Minimap().SetState(statemachine.MINIMAP_OFF)
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			camera.UseCamera().StartAnimation(camera.AnimationOpts{
				Type:        camera.LavaZoneAnimation,
				Duration:    time.Second * 1 / 3,
				Delay:       time.Millisecond * 15,
				MaxRotation: 0.07,
				RotationGap: 0.02,
			})
		}

		pOffset := c.GetScreenCoordsTranslation(p.RawOffset.X, p.RawOffset.Y)

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if pOffset.Y > -sHUD/4 {
				p.SetRawY(p.RawPos.Y - pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y - pSpeed.Y)
				}
				if p.TranslationMovementYBlocked {
					c.MovePosition(0, -pSpeed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			if pOffset.Y < (sAxis.Y * 2) {
				p.SetRawY(p.RawPos.Y + pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y + pSpeed.Y)
				}
				if p.TranslationMovementYBlocked {
					c.MovePosition(0, pSpeed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			p.SetRawX(p.RawPos.X + pSpeed.X)
			if !p.TranslationMovementXBlocked {
				p.SetRawOffsetX(p.RawOffset.X + pSpeed.X)
			}

			if p.TranslationMovementXBlocked {
				c.MovePosition(pSpeed.X, 0)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if pOffset.X > 0 {
				p.SetRawX(p.RawPos.X - pSpeed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - pSpeed.X)
				}

				if p.TranslationMovementXBlocked {
					c.MovePosition(-pSpeed.X, 0)
				}
			}
		}

		if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionUP() {
			p.SetTranslationYMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionDOWN() {
			p.SetTranslationYMovementBlocked(true)
		}

		if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionLEFT() {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionRIGHT() {
			p.SetTranslationXMovementBlocked(true)
		}

		p.UpdateDirection()

		if p.TranslationMovementYBlocked {
			if c.IsLowerZeroCoordY() && p.IsDirectionUP() {
				c.SetZeroPositionY()
				p.SetTranslationYMovementBlocked(false)
			}

			if c.IsOuttaCoordY(wMSize.Y*c.Scale-sAxis.Y*2) && p.IsDirectionDOWN() {
				c.SetPositionY(c.GetWorldCoordY(wMSize.Y*c.Scale - sAxis.Y*2))
				p.SetTranslationYMovementBlocked(false)
			}
		}

		if p.TranslationMovementXBlocked {
			if c.IsLowerZeroCoordX() && p.IsDirectionLEFT() {
				c.SetZeroPositionX()
				p.SetTranslationXMovementBlocked(false)
			}

			if c.IsOuttaCoordX(wMSize.X*c.Scale-sAxis.X*2) && p.IsDirectionRIGHT() {
				c.SetPositionX(c.GetWorldCoordX(wMSize.X*c.Scale - sAxis.X*2))
				p.SetTranslationXMovementBlocked(false)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
