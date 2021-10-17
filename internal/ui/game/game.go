package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/camera"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/world"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	camera.UseCamera().UpdateCamera()

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		w := world.UseWorld()

		opts := &ebiten.DrawImageOptions{}

		screenW, screenH := screen.Size()
		cvx, cvy := camera.UseCamera().GetCameraViewScale(screenW, screenH)
		p := pc.UsePC()

		opts.GeoM.Translate(-p.X, -p.Y)
		opts.GeoM.Scale(cvx, cvy)


		screen.DrawImage(w.Location.Image, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		p := pc.UsePC()
		// w := world.UseWorld()
		m := metadatacollection.GetMetadata("assets/images/heroes/pumpkinhero")
		c := animation.WithAnimation(
			imagecollection.GetImage("assets/images/heroes/pumpkinhero"),
			&m.Animation,
			&p.Equipment.Skin.Animation)

		opts := &ebiten.DrawImageOptions{}

		if history.GetDirection() == direction.LEFT {
			opts.GeoM.Scale(-1, 1)
		}
		if history.GetDirection() == direction.RIGHT {
			opts.GeoM.Scale(1, 1)
		}

		if len(p.Physics.Jump) != 0 {
			select {
			case <-jump.JumpGap.C:
				j := p.Physics.Jump[0]

				if j == direction.UP {
					p.Y -= 2
				}

				if j == direction.DOWN {
					p.Y += 2
				}

				p.Physics.Jump = p.Physics.Jump[1:]
			default:
			}
		}

		screenW, screenH := screen.Size()
		cvx, cvy := camera.UseCamera().GetCameraViewScale(screenW, screenH)
		sx, sy := camera.UseCamera().GetCameraViewSize(screenW, screenH)


		opts.GeoM.Translate((sx + sx/2)/2, (sy + sy/2)/2)
		opts.GeoM.Scale(p.Metadata.Scale.CoefficiantX/cvx, p.Metadata.Scale.CoefficiantY/cvy)

		screen.DrawImage(c, opts)
	})

	// for _, otherC := range pc.PCs{

	// }

	// for _, otherPCs := range {
	// 	img := 	imageloader.Images[players.Equipment.Skin.ImageHash]
	// 	render.SetImageToRender(img, func(i *ebiten.Image) *ebiten.DrawImageOptions {
	// 		return &ebiten.DrawImageOptions{}
	// 	})
	// }
	// screen.DrawImage(, &ebiten.DrawImageOptions{})
	// g.winConf.DrawGameBackground()

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
