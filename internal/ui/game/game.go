package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/physics"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image){
		camera.UseCamera().UpdateCamera(screen)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		w := world.UseWorld()

		opts := &ebiten.DrawImageOptions{}

		// cvx, cvy := camera.UseCamera().GetCameraViewScale(screenW, screenH)
		// ctx, cty := camera.UseCamera().GetCameraViewTranslation(cvx, cvy)		


		// opts.GeoM.Translate(ctx, cty)
		// opts.GeoM.Scale(cvx, cvy)

		opts.GeoM.Concat(camera.UseCamera().MapMatrix)

		screen.DrawImage(w.Location.Image, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		p := pc.UsePC()
		physics.ProcessAnimation(p)
		
		c := animation.WithAnimation(
			p.Image,
			&p.Metadata.Animation,
			&p.Equipment.Skin.Animation)

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Concat(camera.UseCamera().HeroMatrix)

		// opts.GeoM.Scale(p.GetMovementRotation(), 1)

		// tx, ty := camera.UseCamera().GetCharacterTranslation(screen.Size())
		
		// opts.GeoM.Scale(p.Metadata.Scale.CoefficiantX, p.Metadata.Scale.CoefficiantY)
		// opts.GeoM.Translate(tx, ty)

		screen.DrawImage(c, opts)
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
