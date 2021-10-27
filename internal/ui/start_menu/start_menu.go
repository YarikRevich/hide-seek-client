package start_menu

import (
	"fmt"
	"image/color"

	// positioning "github.com/YarikRevich/HideSeek-Client/internal/interface/positioning"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	fontcollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"

	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/background/background")

		opts := &ebiten.DrawImageOptions{}

		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := imagecollection.GetImage("assets/images/system/buttons/settingswheel")
		m := metadatacollection.GetMetadata("assets/images/system/buttons/settingswheel")

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_start")

		opts := &ebiten.DrawImageOptions{}

		fmt.Println(m.Size)

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fontcollection.GetFont("assets/fonts/base")
		p := text.NewPositionSession(
			f, 
			m.Button.Text, 
			m.RawSize.Width,
			m.RawSize.Height, 
			m.Button.TextPosition)

		for p.Next() {
			tx, ty := p.GetPosition()
			ebitentext.Draw(
				img,
				p.GetText(),
				f,
				tx, ty,
				color.White)
		}

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := ebiten.NewImageFromImage(imagecollection.GetImage("assets/images/system/buttons/button"))
		m := metadatacollection.GetMetadata("assets/images/system/buttons/button_join")

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.CoefficiantX, m.Scale.CoefficiantY)

		f := fontcollection.GetFont("assets/fonts/base")
		p := text.NewPositionSession(
			f, 
			m.Button.Text, 
			m.RawSize.Width,
			m.RawSize.Height, 
			m.Button.TextPosition)

		for p.Next() {
			tx, ty := p.GetPosition()
			ebitentext.Draw(
				img,
				p.GetText(),
				f,
				tx, ty,
				color.White)
		}

		screen.DrawImage(img, opts)
	})

	// s.winConf.TextAreas.GameLogo.Clear()
	// IsOld := func(value float64, list []float64) bool {
	// 	//Checks whether new pos is not already used

	// 	for _, i := range list {
	// 		if value == i {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	// max := float64(4)

	// ChangeLogo := func(i float64){
	// 	//It draws logo applying a new color
	// 	//and a new rotation pos appending new
	// 	//params to the 'already used' list

	// 	fmt.Fprint(s.winConf.TextAreas.GameLogo, "Hide&Seek")
	// 	s.winConf.TextAreas.GameLogo.Color = colornames.Orange
	// 	s.winConf.TextAreas.GameLogo.Draw(
	// 		s.winConf.Win,
	// 		pixel.IM.Scaled(s.winConf.TextAreas.GameLogo.Orig, max-i).Rotated(s.winConf.TextAreas.GameLogo.Orig, -.6),
	// 	)
	// 	s.winConf.StartMenu.DrawedTemporally = append(s.winConf.StartMenu.DrawedTemporally, math.Round(i*10)/10)
	// }

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
