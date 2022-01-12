package game

// "fmt"

func Exec() {
	// g := events.UseEvents().Gamepad()
	// k := events.UseEvents().Keyboard()

	// c := world.UseWorld().GetCamera()
	// p := world.UseWorld().GetPC()

	// if g.AreGamepadButtonsCombined(keycodes.GamepadUPButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF1) {
	// 	c.ZoomIn(&p.Base)
	// 	p.UpdateLastActivity()
	// 	return
	// } else if g.AreGamepadButtonsCombined(keycodes.GamepadDOWNButton, keycodes.GamepadLEFTUPPERCLICKERButton) || ebiten.IsKeyPressed(ebiten.KeyF2) {
	// 	c.ZoomOut(&p.Base)
	// 	p.UpdateLastActivity()
	// 	return
	// }

	// poX, poY := p.GetScaledOffsetX(), p.GetScaledOffsetY()
	// if k.IsAnyKeyPressed() {
	// 	p.UpdateLastActivity()
	// 	worldMap := world.UseWorld().GetWorldMap()

	// 	// pWidth, pHeight := p.ModelCombination.Modified.Size.Width, p.ModelCombination.Modified.Size.Height

	// 	p.SaveLastPosition()
	// 	pX, pY := p.GetScaledPosX(), p.GetScaledOffsetY()

	// 	// cX, cY := c.GetScaledPosX(), c.GetScaledPosY()
	// 	cX := c.RawPos.X
	// 	cY := c.RawPos.Y
	// 	// cY := c.GetScaledPosY()

	// 	s := screen.UseScreen()

	// 	hudOffsetHeight := s.GetHUDOffset()
	// 	screenOffsetX := s.GetOffsetX()
	// 	screenOffsetY := s.GetOffsetY()

	// 	// fmt.Println(p.TranslationMovementXBlocked, p.TranslationMovementYBlocked, p.RawOffset, p.RawPos)

	// 	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || g.IsGamepadButtonPressed(keycodes.GamepadUPButton) {
	// 		if pY > hudOffsetHeight {
	// 			if !p.TranslationMovementYBlocked {
	// 				p.SetRawY(p.RawPos.Y - p.Modified.Buffs.Speed.Y)
	// 				p.SetRawOffsetY(p.RawOffset.Y - p.Modified.Buffs.Speed.Y)
	// 			}
	// 		}

	// 		if p.TranslationMovementYBlocked {
	// 			c.SetRawY(c.RawPos.Y - p.Modified.Buffs.Speed.Y)
	// 		}
	// 	}

	// 	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || g.IsGamepadButtonPressed(keycodes.GamepadDOWNButton) {
	// 		if !p.TranslationMovementYBlocked {
	// 			if pY < screenOffsetY*2-(pHeight/1.81) {
	// 				p.SetRawY(p.RawPos.Y + p.Modified.Buffs.Speed.Y)
	// 				p.SetRawOffsetY(p.RawOffset.Y + p.Modified.Buffs.Speed.Y)
	// 			}
	// 		}

	// 		if p.TranslationMovementYBlocked {
	// 			c.SetRawY(c.RawPos.Y + p.Modified.Buffs.Speed.Y)
	// 		}
	// 	}

	// 	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTButton) {
	// 		if !p.TranslationMovementXBlocked {
	// 			if pX < screenOffsetX*2-(pWidth/1.81) {
	// 				p.SetRawX(p.RawPos.X + p.Modified.Buffs.Speed.X)
	// 				p.SetRawOffsetX(p.RawOffset.X + p.Modified.Buffs.Speed.X)
	// 			}
	// 		}

	// 		if p.TranslationMovementXBlocked {
	// 			c.SetRawX(c.RawPos.X + p.Modified.Buffs.Speed.X)
	// 		}
	// 	}

	// 	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.IsGamepadButtonPressed(keycodes.GamepadLEFTButton) {
	// 		if pX > 0 {
	// 			if !p.TranslationMovementXBlocked {
	// 				p.SetRawX(p.RawPos.X - p.Modified.Buffs.Speed.X)
	// 				p.SetRawOffsetX(p.RawOffset.X - p.Modified.Buffs.Speed.X)
	// 			}
	// 		}

	// 		if p.TranslationMovementXBlocked {
	// 			c.SetRawX(c.RawPos.X - p.Modified.Buffs.Speed.X)
	// 		}
	// 	}

	// 	if !p.TranslationMovementXBlocked && s.IsLessAxisXCrossed(poX, p.Modified.Buffs.Speed.X) && p.IsDirectionLEFT() {
	// 		p.SetTranslationXMovementBlocked(true)
	// 	}

	// 	if !p.TranslationMovementXBlocked && s.IsHigherAxisXCrossed(poX, p.Modified.Buffs.Speed.X) && p.IsDirectionRIGHT() {
	// 		p.SetTranslationXMovementBlocked(true)
	// 	}

	// 	if !p.TranslationMovementYBlocked && s.IsLessAxisYCrossed(poY, p.Modified.Buffs.Speed.Y) && p.IsDirectionUP() {
	// 		p.SetTranslationYMovementBlocked(true)
	// 	}

	// 	if !p.TranslationMovementYBlocked && s.IsHigherAxisYCrossed(poY, p.Modified.Buffs.Speed.Y) && p.IsDirectionDOWN() {
	// 		p.SetTranslationYMovementBlocked(true)
	// 	}

	// 	p.UpdateDirection()

	// 	//1.81

	// 	// fmt.Println(cY+screenOffsetY*2, worldMap.ModelCombination.Modified.Size.Height/worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.Y, 1.81, worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.Y)
	// 	// fmt.Println(cY+screenOffsetY*2, worldMap.ModelCombination.Modified.Size.Height/1.81)
	// 	if p.TranslationMovementYBlocked {
	// 		if cY <= -hudOffsetHeight && p.IsDirectionUP() {
	// 			p.SetTranslationYMovementBlocked(false)
	// 		}

	// 		if cY+screenOffsetY*2 >= worldMap.ModelCombination.Modified.Size.Height/1.81 &&
	// 			p.IsDirectionDOWN() {
	// 			p.SetTranslationYMovementBlocked(false)
	// 		}
	// 	}

	// 	if p.TranslationMovementXBlocked {
	// 		if cX+screenOffsetX*2 >= worldMap.ModelCombination.Modified.Size.Width/worldMap.ModelCombination.Modified.RuntimeDefined.ZoomedScale.X*1.81 &&
	// 			p.IsDirectionRIGHT() {
	// 			p.SetTranslationXMovementBlocked(false)
	// 		}
	// 		if cX <= 0 && p.IsDirectionLEFT() {
	// 			p.SetTranslationXMovementBlocked(false)
	// 		}
	// 	}

	// 	if ebiten.IsKeyPressed(ebiten.KeySpace) || g.IsGamepadButtonPressed(keycodes.GamepadRIGHTUPPERCLICKERButton) {
	// 		physics.UsePhysics().Jump().Calculate()
	// 	}

	// 	//
	// }
}
