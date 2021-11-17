package game

import (
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
	wsx, wsy := w.GetZoomedMapScale()
	wm := w.GetMetadata().Modified

	p := objects.UseObjects().PC()
	zx, zy := p.GetZoomedRawPos(w.GetZoomedMapScale())
	pmm := p.GetMetadata().Modified
	pmo := p.GetMetadata().Origin

	if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
		camera.UseCamera().ZoomIn()
		return
	}

	if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
		camera.UseCamera().ZoomOut()
		return
	}

	

	if k.IsAnyKeyPressed() {
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
			if p.GetY() > 0 {
				p.SetRawY(p.RawPos.Y - pmm.Buffs.Speed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawPosForCameraY(p.RawPosForCamera.Y - pmm.Buffs.Speed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			if zy+pmm.Buffs.Speed.Y+pmo.Size.Height < wm.Size.Height*wsy {
				p.SetRawY(p.RawPos.Y + pmm.Buffs.Speed.Y)
				if !p.TranslationMovementYBlocked {
					p.SetRawPosForCameraY(p.RawPosForCamera.Y + pmm.Buffs.Speed.Y)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			if zx+pmm.Buffs.Speed.X+pmo.Size.Width < wm.Size.Width*wsx {
				p.SetRawX(p.RawPos.X + pmm.Buffs.Speed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawPosForCameraX(p.RawPosForCamera.X + pmm.Buffs.Speed.X)
				}
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if p.GetX() > 0 {
				p.SetRawX(p.RawPos.X - pmm.Buffs.Speed.X)
				if !p.TranslationMovementXBlocked {
					p.SetRawPosForCameraX(p.RawPosForCamera.X - pmm.Buffs.Speed.X)
				}
			}
		}

		p.UpdateDirection()

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
