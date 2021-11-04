package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/physics"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	w := objects.UseObjects().World()
	p := objects.UseObjects().PC()
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		p.SaveLastPosition()
		camera.UseCamera().UpdateCamera()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		opts := &ebiten.DrawImageOptions{}
		// opts.GeoM.Concat(camera.UseCamera().MapMatrix)

		screen.DrawImage(w.GetImage(), opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		// physics.ProcessAnimation(&p.Object)
		c := animation.WithAnimation(&p.Object)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Concat(camera.UseCamera().HeroMatrix)

		screen.DrawImage(c, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		for _, v := range w.PCs {
			// if v.ID != p.ID {
			opts := &ebiten.DrawImageOptions{}
			screen.DrawImage(v.GetImage(), opts)
			// }
		}
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		for _, v := range w.Weapons {
			opts := &ebiten.DrawImageOptions{}
			screen.DrawImage(v.GetImage(), opts)
		}
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		for _, v := range w.Ammo {
			opts := &ebiten.DrawImageOptions{}
			screen.DrawImage(v.GetImage(), opts)
		}
	})

	// // g.winConf.DrawGoldChest()

	// g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	// Animation.NewDefaultSwordAnimator(g.winConf, g.userConfig).Move()
	// Animation.NewIconAnimator(g.winConf, g.userConfig).Move()

	// for _, value := range g.winConf.GameProcess.OtherUsers {
	// 	Animation.NewDefaultSwordAnimator(g.winConf, value).Move()
	// 	Animation.NewIconAnimator(g.winConf, value).Move()
	// }

	// g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.Pos.X)*2.5)-31, (float64(g.userConfig.Pos.Y)*2.5)-30))

	// g.winConf.DrawElementsPanel()

	// g.mapComponents.GetCam().UpdateCam()

	// var bias float64
	// for i := 0; i <= g.userConfig.GameInfo.Health; i++ {
	// 	g.winConf.DrawHPHeart(
	// 		pixel.V(-40+bias, 1200),
	// 	)
	// 	bias += 100
	// }

	// g.winConf.DrawWeaponIcon(g.userConfig.GameInfo.WeaponName)

	// if g.userConfig.GameInfo.Health < 1 {
	// 	g.mapComponents.GetCam().SetDefaultCam()
	// 	g.currState.MainStates.SetStartMenu()
	// }
}
