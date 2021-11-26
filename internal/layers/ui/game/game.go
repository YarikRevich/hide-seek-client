package game

import (
	// "github.com/x/color"

	// "image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	worldMap := world.UseWorld().GetWorldMap()

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := worldMap.GetImage()

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(worldMap.ModelCombination.Modified.ZoomedScale.X, worldMap.ModelCombination.Modified.ZoomedScale.Y)

		c := world.UseWorld().GetCamera()
		opts.GeoM.Translate(-c.GetScaledPosX(), -c.GetScaledPosY())

		screen.DrawImage(img, opts)
	})

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	screenWidth, screenHeight := screen.Size()

	// 	i := ebiten.NewImage(screenWidth, screenHeight/10)

	// 	i.Fill(color.Black)
	// 	screen.DrawImage(i, nil)
	// })

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		p := world.UseWorld().GetPC()
		pcs := world.UseWorld().GetPCs()

		for _, pc := range pcs {
			img := pc.GetAnimatedImage()

			opts := &ebiten.DrawImageOptions{}

			opts.GeoM.Scale(pc.GetMovementRotation(), 1)
			opts.GeoM.Scale(pc.ModelCombination.Modified.ZoomedScale.X,
				pc.ModelCombination.Modified.ZoomedScale.Y)

			if pc.IsEqualTo(p.Base) {
				pc.GetScaledPosX()
				opts.GeoM.Translate(pc.GetScaledOffsetX(), pc.GetScaledOffsetX())
			} else {
				opts.GeoM.Translate(pc.GetScaledPosX(), pc.GetScaledPosY())
			}

			screen.DrawImage(img, opts)
		}
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("hud/health/health")
		mm := sources.UseSources().Metadata().GetMetadata("hud/health/health").Modified

		p := world.UseWorld().GetPC()
		for i := 0; i < int(p.Health); i++ {
			opts := &ebiten.DrawImageOptions{}

			opts.GeoM.Translate(mm.Margins.LeftMargin+(mm.Size.Width+(10/mm.Scale.X))*float64(i), mm.Margins.TopMargin)
			opts.GeoM.Scale(mm.Scale.X, mm.Scale.Y)

			screen.DrawImage(img, opts)
		}
	})

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	weapons := world.UseWorld().GetWeapons()

	// 	for _, weapon := range weapons {
	// 		opts := &ebiten.DrawImageOptions{}
	// 		screen.DrawImage(weapon.GetImage(), opts)
	// 	}
	// })

	// render.UseRender().SetToRender(func(screen *ebiten.Image) {
	// 	ammos := world.UseWorld().GetAmmos()

	// 	for _, ammo := range ammos {
	// 		opts := &ebiten.DrawImageOptions{}
	// 		screen.DrawImage(ammo.GetImage(), opts)
	// 	}
	// })

	// // // g.winConf.DrawGoldChest()

	// // g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	// // Animation.NewDefaultSwordAnimator(g.winConf, g.userConfig).Move()
	// // Animation.NewIconAnimator(g.winConf, g.userConfig).Move()

	// // for _, value := range g.winConf.GameProcess.OtherUsers {
	// // 	Animation.NewDefaultSwordAnimator(g.winConf, value).Move()
	// // 	Animation.NewIconAnimator(g.winConf, value).Move()
	// // }

	// // g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.Pos.X)*2.5)-31, (float64(g.userConfig.Pos.Y)*2.5)-30))

	// // g.winConf.DrawElementsPanel()

	// // g.mapComponents.GetCam().UpdateCam()

	// // var bias float64
	// // for i := 0; i <= g.userConfig.GameInfo.Health; i++ {
	// // 	g.winConf.DrawHPHeart(
	// // 		pixel.V(-40+bias, 1200),
	// // 	)
	// // 	bias += 100
	// // }

	// // g.winConf.DrawWeaponIcon(g.userConfig.GameInfo.WeaponName)

	// // if g.userConfig.GameInfo.Health < 1 {
	// // 	g.mapComponents.GetCam().SetDefaultCam()
	// // 	g.currState.MainStates.SetStartMenu()
	// // }
}
