package game

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/camera"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	worldMap := world.UseWorld().GetWorldMap()
	c := camera.UseCamera()
	s := screen.UseScreen()

	if statemachine.UseStateMachine().Minimap().GetState() == statemachine.MINIMAP_ON {
		render.UseRender().SetToRender(func(screen *ebiten.Image) {

		})
	}

	render.UseRender().SetToRender(func(i *ebiten.Image) {
		img := worldMap.GetImage()
		// scale := worldMap.GetScale()

		opts := &ebiten.DrawImageOptions{}
		// s := screen.UseScreen().GetSize()
		// opts.GeoM.Scale((c.Scale.X*s.X)/100, (c.Scale.Y*s.Y)/100)
		// opts.GeoM.Scale(scale.X*c.Scale.X, scale.Y*c.Scale.Y)
		opts.GeoM.Concat(c.GetCameraOptions().GeoM)
		// opts.GeoM.Scale(scale.X, scale.Y)
		// cPos := c.GetCameraTranslation()
		// opts.GeoM.Translate(cPos.X/(c.Scale.X*s.X)/100, cPos.Y/(c.Scale.Y*s.Y)/100)
		// cPos := c.GetCameraTranslation()
		// opts.GeoM.Translate(cPos.X*scale.X, cPos.Y*scale.Y)

		if statemachine.UseStateMachine().Minimap().GetState() == statemachine.MINIMAP_ON {
			opts.Filter = ebiten.FilterLinear
		}
		i.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(i *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("maps/helloween/elements/torch")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(100, 100)
		opts.GeoM.Translate(0, s.GetHUDOffset())
		i.DrawImage(img, opts)
	})

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	if statemachine.UseStateMachine().PCs().GetState(p.ID) == statemachine.PC_DEAD_NOW {
	// 		latency.UseLatency().Timings().ExecFor(func() {
	// 		}, func() {
	// 			statemachine.UseStateMachine().PCs().SetState(p.ID, statemachine.PC_ALIVE)
	// 		}, statemachine.UI_GAME, time.Second)
	// 	}
	// })

	render.UseRender().SetToRender(func(i *ebiten.Image) {
		s := screen.UseScreen()
		hudHeight := s.GetHUDOffset()
		img := ebiten.NewImage(int(s.GetSize().X), int(hudHeight))

		opts := &ebiten.DrawImageOptions{}

		img.Fill(color.Black)
		i.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {

		pcs := world.UseWorld().GetPCs()

		for _, pc := range pcs {
			img := pc.GetAnimatedImage()

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(pc.GetMovementRotation(), 1)

			pcScale := pc.GetScale()

			opts.GeoM.Scale(pcScale.X, pcScale.Y)
			opts.GeoM.Scale(c.Scale.X/2, c.Scale.Y/2)

			if pc.IsEqualTo(&pc.Base) {
				pOffset := c.GetScreenCoordsTranslation(pc.RawOffset.X, pc.RawOffset.Y)
				opts.GeoM.Translate(pOffset.X, pOffset.Y)
			} else {
				pPos := c.GetScreenCoordsTranslation(pc.RawPos.X, pc.RawPos.Y)
				opts.GeoM.Translate(pPos.X, pPos.Y)
			}

			// opts.GeoM.Translate(0, s.GetHUDOffset())

			screen.DrawImage(img, opts)
		}
	})

	// 	// render.UseRender().SetToRender(func(i *ebiten.Image) {
	// 	// 	img := p.GetImage()
	// 	// 	opts := &ebiten.DrawImageOptions{}
	// 	// 	opts.GeoM.Scale(2, 2)
	// 	// 	opts.GeoM.Translate((p.GetScaledOffsetX() - c.AlignOffset.X), p.GetScaledOffsetY()-c.AlignOffset.Y)
	// 	// 	// opts.GeoM.Translate((pc.GetScaledOffsetX()-cam.AlignOffset.X), p.Position[1]+(pc.GetScaledOffsetY()-cam.AlignOffset.Y))
	// 	// 	// opts.GeoM.Rotate(p.Rotation)
	// 	// 	opts.ColorM.Apply(color.White)

	// 	// life := float32(p.LifeRemaining / p.LifeTime)

	// 	// color := glm.QuatLerp(&p.ColorEnd, &p.ColorBegin, life)
	// 	// // opts.ColorM.Translate(float64(color.X()), float64(color.Y()), float64(color.Z()), float64(color.W))
	// 	// opts.ColorM.Apply(colornames.Black)

	// 	// scale := glm.QuatLerp(&p.SizeEnd, &p.SizeBegin, life)
	// 	// opts.GeoM.Scale(float64(scale.X()), float64(scale.X()))

	// 	// fmt.Println("SCALE: ", scale.X(), "TRANSLATE: ", p.Position[0], p.Position[1], "COLOR: ", float64(color.X()), float64(color.Y()), float64(color.Z()), float64(color.W))
	// 	// 	fmt.Println(opts)

	// 	// 	i.DrawImage(img, opts)
	// 	// })
	// 	render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 		img := sources.UseSources().Images().GetCopyOfImage("hud/health/health")
	// 		mm := sources.UseSources().Metadata().GetMetadata("hud/health/health").Modified

	// 		p := world.UseWorld().GetPC()
	// 		for i := 0; i < int(p.Health); i++ {
	// 			opts := &ebiten.DrawImageOptions{}

	// 			opts.GeoM.Translate(mm.Margins.LeftMargin+(mm.Size.Width+(10/mm.Scale.X))*float64(i), mm.Margins.TopMargin)
	// 			opts.GeoM.Scale(mm.Scale.X, mm.Scale.Y)

	// 			screen.DrawImage(img, opts)
	// 		}
	// 	})

	// 	// // render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	// // 	weapons := world.UseWorld().GetWeapons()

	// 	// // 	for _, weapon := range weapons {
	// 	// // 		opts := &ebiten.DrawImageOptions{}
	// 	// // 		screen.DrawImage(weapon.GetImage(), opts)
	// 	// // 	}
	// 	// // })

	// 	// // render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	// // 	ammos := world.UseWorld().GetAmmos()

	// 	// // 	for _, ammo := range ammos {
	// 	// // 		opts := &ebiten.DrawImageOptions{}
	// 	// // 		screen.DrawImage(ammo.GetImage(), opts)
	// 	// // 	}
	// 	// // })

	// 	// // // // g.winConf.DrawGoldChest()

	// 	// // // g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	// 	// // // Animation.NewDefaultSwordAnimator(g.winConf, g.userConfig).Move()
	// 	// // // Animation.NewIconAnimator(g.winConf, g.userConfig).Move()

	// 	// // // for _, value := range g.winConf.GameProcess.OtherUsers {
	// 	// // // 	Animation.NewDefaultSwordAnimator(g.winConf, value).Move()
	// 	// // // 	Animation.NewIconAnimator(g.winConf, value).Move()
	// 	// // // }

	// 	// // // g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.Pos.X)*2.5)-31, (float64(g.userConfig.Pos.Y)*2.5)-30))

	// 	// // // g.winConf.DrawElementsPanel()

	// 	// // // g.mapComponents.GetCam().UpdateCam()

	// 	// // // var bias float64
	// 	// // // for i := 0; i <= g.userConfig.GameInfo.Health; i++ {
	// 	// // // 	g.winConf.DrawHPHeart(
	// 	// // // 		pixel.V(-40+bias, 1200),
	// 	// // // 	)
	// 	// // // 	bias += 100
	// 	// // // }

	// 	// // // g.winConf.DrawWeaponIcon(g.userConfig.GameInfo.WeaponName)

	// 	// // // if g.userConfig.GameInfo.Health < 1 {
	// 	// // // 	g.mapComponents.GetCam().SetDefaultCam()
	// 	// // // 	g.currState.MainStates.SetStartMenu()
	// 	// // // }
}
