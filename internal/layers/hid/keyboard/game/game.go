package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/keycodes"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/core/physics"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec() {
	g := events.UseEvents().Gamepad()
	k := events.UseEvents().Keyboard()

	// w := objects.UseObjects().World()
	// msw, msh := camera.UseCamera().GetMapScale()
	p := objects.UseObjects().PC()

	// fmt.Println(sw, w.Metadata.Size.Width, "KEYBOARD GAME")

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
			if p.RawPos.Y > 0 {
				p.SetY(p.RawPos.Y - p.Metadata.Buffs.Speed.Y)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
			// if p.RawPos.Y+p.Metadata.Buffs.Speed.Y < w.Metadata.Size.Height * msh {
				p.SetY(p.RawPos.Y + p.Metadata.Buffs.Speed.Y)
			// } else {
				// p.SetY(p.RawPos.Y + (w.Metadata.Size.Height - p.RawPos.Y))
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
			
			// if p.RawPos.X+p.Metadata.Buffs.Speed.X < w.Metadata.Size.Width * msw {
				p.SetX(p.RawPos.X + p.Metadata.Buffs.Speed.X)
			// } else {
				// p.SetX(p.RawPos.X + (w.Metadata.Size.Width - p.RawPos.X))
			// }
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
			if p.RawPos.X > 0 {

				p.SetX(p.RawPos.X - p.Metadata.Buffs.Speed.X)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
			physics.UsePhysics().Jump().Calculate()
		}
	}
}
