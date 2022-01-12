package settingsmenu

import (
	"image/color"

	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/text/positioning"

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
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/back")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(s.X, s.Y)

		screen.DrawImage(img, opts)
	})

	f := sources.UseSources().Font().GetFont("base")

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/inputs/input")
		m := sources.UseSources().Metadata().GetMetadata("system/inputs/settingsmenuinput")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(s.X, s.Y)

		t := events.UseEvents().Input().SettingsMenuNameBuffer.Read()

		p := positioning.UsePositioning().Input()
		p.Init(img, m, f, t)
		p.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/buttons/button")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/button_save_config")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(s.X, s.Y)

		p := positioning.UsePositioning().Button()
		p.Init(img, m, f, m.Text.Symbols)
		p.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		var (
			img *ebiten.Image
			m   *sources.MetadataModel
		)
		ms := m.GetMargins()

		mt := sources.UseSources().Metadata().GetMetadata("fonts/settingsmenu/settingsmenu")
		mts := mt.GetMargins()
		switch statemachine.UseStateMachine().SettingsMenuCheckbox().GetState() {
		case statemachine.UI_SETTINGS_MENU_CHECKBOX_OFF:
			img = sources.UseSources().Images().GetImage("system/checkbox/greencheckboxoff")
			m = sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxoff")
			text.Draw(screen, "Enable LAN server", f, int(mts.X), int(mts.Y), color.White)
		case statemachine.UI_SETTINGS_MENU_CHECKBOX_ON:
			img = sources.UseSources().Images().GetImage("system/checkbox/greencheckboxon")
			m = sources.UseSources().Metadata().GetMetadata("system/checkbox/greencheckboxon")
			text.Draw(screen, "Disable LAN server", f, int(mts.X), int(mts.Y), color.White)
		}

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(m.Scale.X, m.Scale.Y)

		screen.DrawImage(img, opts)
	})
}
