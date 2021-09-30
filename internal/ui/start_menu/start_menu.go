package start_menu

import (
	// "fmt"

	// "github.com/YarikRevich/HideSeek-Client/internal/connection"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	back := imageloader.GetImage("/images/menues/background/StartMenu")

	render.SetImageToRender(render.Cell{Image: back, CallBack: func(screen *ebiten.Image) *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := back.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		return opts
	}})

	setting := imageloader.GetImage("/images/menues/panels/settingswheel")
	render.SetImageToRender(render.Cell{Image: setting, CallBack: func(screen *ebiten.Image) *ebiten.DrawImageOptions {
		return new(ebiten.DrawImageOptions)
	}})

	// screen.DrawImage(img, opts)
	// 	s.winConf.TextAreas.GameLogo.Clear()
	// 	IsOld := func(value float64, list []float64) bool {
	// 		//Checks whether new pos is not already used

	// 		for _, i := range list {
	// 			if value == i {
	// 				return true
	// 			}
	// 		}
	// 		return false
	// 	}

	// 	max := float64(4)

	// 	ChangeLogo := func(i float64){
	// 		//It draws logo applying a new color
	// 		//and a new rotation pos appending new
	// 		//params to the 'already used' list

	// 		fmt.Fprint(s.winConf.TextAreas.GameLogo, "Hide&Seek")
	// 		s.winConf.TextAreas.GameLogo.Color = colornames.Orange
	// 		s.winConf.TextAreas.GameLogo.Draw(
	// 			s.winConf.Win,
	// 			pixel.IM.Scaled(s.winConf.TextAreas.GameLogo.Orig, max-i).Rotated(s.winConf.TextAreas.GameLogo.Orig, -.6),
	// 		)
	// 		s.winConf.StartMenu.DrawedTemporally = append(s.winConf.StartMenu.DrawedTemporally, math.Round(i*10)/10)
	// 	}

	// 	switch s.winConf.StartMenu.Regime {
	// 	case 0:
	// 		//This regime makes logo smaller

	// 		for i := max; i != 0; i = i - 0.1 {
	// 			if !IsOld(math.Round(i*10)/10, s.winConf.StartMenu.DrawedTemporally){
	// 				ChangeLogo(i)
	// 				s.winConf.TextAreas.GameLogo.Orig = pixel.V(s.winConf.TextAreas.GameLogo.Orig.X-math.Round(i*10)/10, s.winConf.TextAreas.GameLogo.Orig.Y+math.Round(i*10)/10)
	// 				break
	// 			}
	// 		}
	// 		if len(s.winConf.StartMenu.DrawedTemporally) == int(max*10+1) {
	// 			s.winConf.StartMenu.DrawedTemporally = []float64{}
	// 			s.winConf.StartMenu.Regime = 1
	// 		}
	// 	case 1:
	// 		//This regime makes logo bigger

	// 		for i := 0.0; i != max; i = i + 0.1 {
	// 			if !IsOld(math.Round(i*10)/10, s.winConf.StartMenu.DrawedTemporally) {
	// 				ChangeLogo(i)
	// 				s.winConf.TextAreas.GameLogo.Orig = pixel.V(s.winConf.TextAreas.GameLogo.Orig.X+math.Round(i*10)/10, s.winConf.TextAreas.GameLogo.Orig.Y-math.Round(i*10)/10)
	// 				break
	// 			}
	// 		}
	// 		if len(s.winConf.StartMenu.DrawedTemporally) == int(max*10+1) {
	// 			s.winConf.StartMenu.DrawedTemporally = []float64{}
	// 			s.winConf.StartMenu.Regime = 0
	// 		}
	// 	}

}
