package game

import (
	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	//

	// "github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/player_mechanics/animation"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	back := imageloader.GetImage("/images/maps/default/background/Game")

	render.SetImageToRender(render.Cell{Image: back, CallBack: func(i *ebiten.Image) *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := back.Size()
		screenW, screenH := i.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		return opts
	}})

	p := pc.GetPC()
	c := animation.WithAnimation(
		imageloader.GetImage("/images/heroes/pumpkinhero"),
		metadataloader.Metadata["/images/heroes/pumpkinhero"],
		&p.Equipment.Skin.Animation)
	render.SetImageToRender(render.Cell{Image: c, CallBack: func(i *ebiten.Image) *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}

		if history.GetDirection() == history.LEFT {
			opts.GeoM.Scale(-1, 1)
		}
		if history.GetDirection() == history.RIGHT {
			opts.GeoM.Scale(1, 1)
		}

		opts.GeoM.Translate(p.X, p.Y)
		return opts
	}})

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
