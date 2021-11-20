package settingsmenu

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text/positioning"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/background/background")

		opts := &ebiten.DrawImageOptions{}
		imageW, imageH := img.Size()
		screenW, screenH := screen.Size()
		opts.GeoM.Scale(float64(screenW)/float64(imageW), float64(screenH)/float64(imageH))

		screen.DrawImage(img, opts)
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetImage("system/buttons/back")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back").Modified

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

		screen.DrawImage(img, opts)
	})

	f := sources.UseSources().Font().GetFont("base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/inputs/input")
		mm := sources.UseSources().Metadata().GetMetadata("system/inputs/settingsmenuinput").Modified
		mo := sources.UseSources().Metadata().GetMetadata("system/inputs/settingsmenuinput").Origin

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(mm.Margins.LeftMargin, mm.Margins.TopMargin)
		opts.GeoM.Scale(mm.Scale.X, mm.Scale.Y)

		t := events.UseEvents().Input().SettingsMenuNameBuffer.Read()

		s := positioning.UsePositioning().Input()
		s.Init(img, mo, f, t)
		s.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/buttons/button")
		mm := sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config").Modified
		mo := sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config").Origin

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(mm.Margins.LeftMargin, mm.Margins.TopMargin)
		opts.GeoM.Scale(mm.Scale.X, mm.Scale.Y)

		s := positioning.UsePositioning().Button()
		s.Init(img, mo, f, mo.Text.Symbols)
		s.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		var (
			img *ebiten.Image
			m   *sources.Model
		)

		mt := sources.UseSources().Metadata().GetMetadata("fonts/settingsmenu/settingsmenu").Modified
		switch statemachine.UseStateMachine().SettingsMenuCheckbox().GetState() {
		case statemachine.UI_SETTINGS_MENU_CHECKBOX_OFF:
			img = sources.UseSources().Images().GetImage("system/checkbox/greencheckboxoff")
			m = sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxoff").Modified
			text.Draw(screen, "Enable LAN server", f, int(mt.Margins.LeftMargin), int(mt.Margins.TopMargin), color.White)
		case statemachine.UI_SETTINGS_MENU_CHECKBOX_ON:
			img = sources.UseSources().Images().GetImage("system/checkbox/greencheckboxon")
			m = sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxon").Modified
			text.Draw(screen, "Disable LAN server", f, int(mt.Margins.LeftMargin), int(mt.Margins.TopMargin), color.White)
		}

		

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(m.Margins.LeftMargin, m.Margins.TopMargin)
		opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

		screen.DrawImage(img, opts)
	})
}
