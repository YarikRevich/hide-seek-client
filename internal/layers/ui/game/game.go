package game

import (
	"fmt"
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

// import (
// 	"image/color"
// 	"time"

// 	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/render"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
// 	"github.com/YarikRevich/hide-seek-client/internal/core/world"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

func Draw() {
	worldMap := world.UseWorld().GetWorldMap()
	c := world.UseWorld().GetCamera()
	p := world.UseWorld().GetPC()
	s := screen.UseScreen()
	sAxis := s.GetAxis()
	// size := s.GetSize()
	// lastSize := s.GetLastSize()
	// fmt.Println(world.UseWorld().GetGameSettings().IsGameStarted)

	if statemachine.UseStateMachine().Minimap().GetState() == statemachine.MINIMAP_ON {
		render.UseRender().SetToRender(func(screen *ebiten.Image) {

		})
	}

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := worldMap.GetImage()

		opts := &ebiten.DrawImageOptions{}
		wMScale := c.GetZoomedScale(&worldMap.Base)
		opts.GeoM.Scale(wMScale.X, wMScale.Y)
		// fmt.Println(c.DeltaOffset)

		pScale := c.GetZoomedScale(&p.Base)

		fmt.Println(((c.RawPos.Y-sAxis.Y)*pScale.Y+sAxis.Y)+sAxis.Y*2, c.RawPos.Y, sAxis.Y, pScale.Y, "CAMERA")
		fmt.Println((c.RawPos.Y * pScale.Y))
		// fmt.Println(-((c.RawPos.X+sAxisXHalf)*pScale.X - sAxisXHalf), -((c.RawPos.Y+sAxisYHalf)*pScale.Y - sAxisYHalf))
		// opts.GeoM.Translate(-(c.RawPos.X * pScale.X), -(c.RawPos.Y * pScale.Y))
		opts.GeoM.Translate(-((c.RawPos.X+sAxis.X)*pScale.X - sAxis.X), -((c.RawPos.Y+sAxis.Y)*pScale.Y - sAxis.Y))

		opts.GeoM.Translate(0, s.GetHUDOffset())

		if statemachine.UseStateMachine().Minimap().GetState() == statemachine.MINIMAP_ON {
			opts.Filter = ebiten.FilterLinear
		}
		screen.DrawImage(img, opts)
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

			pScale := c.GetZoomedScale(&pc.Base)
			opts.GeoM.Scale(pScale.X, pScale.Y)
			// opts.GeoM.Scale(c.Zoom/100, c.Zoom/100)

			if pc.IsEqualTo(&pc.Base) {
				// fmt.Println(((pc.RawOffset.Y+sAxis.Y)*pScale.Y - sAxis.Y), "PC")
				opts.GeoM.Translate(((pc.RawOffset.X) * pScale.X), ((pc.RawOffset.Y-sAxis.Y)*pScale.Y + sAxis.Y))
			} else {
				opts.GeoM.Translate(pc.RawOffset.X, pc.RawOffset.Y)
			}

			opts.GeoM.Translate(0, s.GetHUDOffset())

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
