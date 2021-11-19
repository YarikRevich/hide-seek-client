package game

import (
	// "fmt"

	// "fmt"

	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/physics"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()

	w := objects.UseObjects().World()
	p := objects.UseObjects().PC()
	c := objects.UseObjects().Camera()

	pmm := p.GetMetadata().Modified
	pmo := p.GetMetadata().Origin

	wm := w.GetMetadata().Modified
	mapScaleX, mapScaleY := w.GetZoomedMapScale()


	pZoomedX, pZoomedY := p.GetZoomedRawPos(w.GetZoomedMapScale())
	pZoomedOffsetX, _ := p.GetZoomedRawPosForCamera(w.GetZoomedMapScale())

	cZoomedX, cZoomedY := c.GetZoomedRawPos(w.GetZoomedMapScale())

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		camera.UseCamera().ZoomIn()
		return
	}

	if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		camera.UseCamera().ZoomOut()
		return
	}

	if k.IsAnyKeyPressed() {

		pZoomedSpeedX, pZoomedSpeedY := pmm.Buffs.Speed.X*mapScaleX, pmm.Buffs.Speed.Y*mapScaleY
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {

			if pZoomedY > 0 {
				p.SetRawY(p.RawPos.Y - pZoomedSpeedY)
				if !p.TranslationMovementYBlocked {
					p.SetRawPosForCameraY(p.RawPosForCamera.Y - pZoomedSpeedY)
				}
			}

			if cZoomedY > 0 {
				c.SetRawY(c.RawPos.Y - pZoomedSpeedY)
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

			fmt.Println(pZoomedSpeedX)
			p.SetRawX(p.RawPos.X + pZoomedSpeedX)
			if !p.TranslationMovementXBlocked {
				p.SetRawPosForCameraX(p.RawPosForCamera.X + pZoomedSpeedX)
			}
			// }

			// fmt.Println(cZoomedX, wm.Size.Width*wsx, cZoomedX < wm.Size.Width*wsx)
			// if c.GetRawX() < wm.Size.Width*wsx {
			fmt.Println(cZoomedX, pZoomedOffsetX*2, pmo.Size.Width, pmm.Buffs.Speed.X < wm.Size.Width*mapScaleX, "PUK")
			if p.TranslationMovementXBlocked {
				c.SetRawX(c.RawPos.X + pZoomedSpeedX)
			}
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {

			if pZoomedX > 0 {
				p.SetRawX(p.RawPos.X - pZoomedSpeedX)
				if !p.TranslationMovementXBlocked {
					p.SetRawPosForCameraX(p.RawPosForCamera.X - pZoomedSpeedX)
				}
			}

			if cZoomedX > 0 {
				if p.TranslationMovementXBlocked {
					c.SetRawX(c.RawPos.X - pZoomedSpeedX)
				}
			}
		}

		p.UpdateDirection()

		if !p.TranslationMovementXBlocked && w.IsAxisXCrossedBy(p) {
			p.SetTranslationXMovementBlocked(true)
		}

		if !p.TranslationMovementYBlocked && w.IsAxisYCrossedBy(p) {
			p.SetTranslationYMovementBlocked(true)
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
