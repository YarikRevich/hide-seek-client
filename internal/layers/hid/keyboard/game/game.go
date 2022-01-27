package game

import (
	"fmt"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/keycodes"
	"github.com/YarikRevich/hide-seek-client/internal/core/physics"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
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

	// p.RawOffset := p.MetadataModel.GetOffset()
	pSpeed := p.MetadataModel.GetBuffSpeed()
	// pScale := c.GetZoomedScale(&p.Base)
	wM := world.UseWorld().GetWorldMap()
	// wMScale := c.GetZoomedScale(&wM.Base)
	// wMScale := c.GetZoomedScale(&wM.Base)
	wMSize := wM.GetSize()

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		// c.ZoomIn(&p.Base)
		camera.Cam.Zoom(1.1)
		p.UpdateLastActivity()
		return
	} else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		// c.ZoomOut(&p.Base)
		camera.Cam.Zoom(0.9)
		p.UpdateLastActivity()
		return
	}

	if k.IsAnyKeyPressed() || g.IsAnyButtonPressed() {
		p.UpdateLastActivity()
		p.UpdateLastPosition()

		s := screen.UseScreen()
		sAxis := s.GetAxis()
		sHUD := s.GetHUDOffset()

		// fmt.Println(pSpeed)

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
			camera.Cam.RunShakingLimitedAnimation(time.Second*1/3, time.Millisecond*15, 0.07, 0.02, make(chan int))
		}

		pScreenOffsetX, pScreenOffsetY := camera.Cam.GetScreenCoordsTranslation(p.RawOffset.X, p.RawOffset.Y)
		pOffset := types.Vec2{X: pScreenOffsetX, Y: pScreenOffsetY}
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if pOffset.Y > -sHUD/4 {
				p.SetRawY(p.RawPos.Y - pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y - pSpeed.Y)
				}
				if p.TranslationMovementYBlocked {
					camera.Cam.MovePosition(0, -pSpeed.Y)
					c.SetRawY(c.RawPos.Y - pSpeed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			// fmt.Println(pOffset.Y, ((sAxis.Y * 2) - sHUD - p.MetadataModel.GetSize().X))
			// -sHUD*2 - p.MetadataModel.GetSize().X
			if pOffset.Y < (sAxis.Y * 2) {
				p.SetRawY(p.RawPos.Y + pSpeed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawOffsetY(p.RawOffset.Y + pSpeed.Y)
				}
				if p.TranslationMovementYBlocked {
					camera.Cam.MovePosition(0, pSpeed.Y)
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
				camera.Cam.MovePosition(pSpeed.X, 0)
				c.SetRawX(c.RawPos.X + pSpeed.X)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if pOffset.X > 0 {
				p.SetRawX(p.RawPos.X - pSpeed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawOffsetX(p.RawOffset.X - pSpeed.X)
				}

				if p.TranslationMovementXBlocked {
					camera.Cam.MovePosition(-pSpeed.X, 0)
					c.SetRawX(c.RawPos.X - pSpeed.X)
				}
			}
		}

		// fmt.Println(!p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionUP())
		if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionUP() {
			p.SetTranslationYMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(pOffset.Y, pSpeed.Y) && p.IsDirectionDOWN() {
			p.SetTranslationYMovementBlocked(true)
		}

		// fmt.Println(pOffset.X)
		if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionLEFT() {
			p.SetTranslationXMovementBlocked(true)
		}

		// fmt.Println(p.RawOffset.X*pScale.X, pSpeed.X, s.IsHigherAxisXCrossed(p.RawOffset.X*pScale.X, pSpeed.X))
		if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(pOffset.X, pSpeed.X) && p.IsDirectionRIGHT() {
			p.SetTranslationXMovementBlocked(true)
		}

		p.UpdateDirection()

		cScreenOffsetX, cScreenOffsetY := camera.Cam.GetCameraTranslation()
		cPos := types.Vec2{X: cScreenOffsetX, Y: cScreenOffsetY}

		if p.TranslationMovementYBlocked {
			if -cPos.Y <= 0 && p.IsDirectionUP() {
				camera.Cam.SetPositionY(0)
				p.SetTranslationYMovementBlocked(false)
			}

			fmt.Println(wMSize.Y-sAxis.Y, camera.Cam.Y, sHUD)
			if c.IsOuttaRange(wMSize.Y*camera.Cam.Scale-sAxis.Y+sHUD/2, -cPos.Y) {
				camera.Cam.SetPositionY(wMSize.Y - sAxis.Y + sHUD/2)
				p.SetTranslationYMovementBlocked(false)
			}
		}

		if p.TranslationMovementXBlocked {
			if -cPos.X <= 0 && p.IsDirectionLEFT() {
				camera.Cam.SetPositionX(0)
				p.SetTranslationXMovementBlocked(false)
			}

			if c.IsOuttaRange(wMSize.X*camera.Cam.Scale-sAxis.X*2, -cPos.X) {
				camera.Cam.SetPositionX(wMSize.X - sAxis.X)
				p.SetTranslationXMovementBlocked(false)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
