package joinmenu

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/events"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/text/positioning"
	"github.com/hajimehoshi/ebiten/v2"
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
		m := sources.UseSources().Metadata().GetMetadata("system/inputs/joingameinput")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(s.X, s.Y)

		t := events.UseEvents().Input().JoinGameBuffer.Read()

		p := positioning.UsePositioning().Input()
		p.Init(img, m, f, t)
		p.Draw()

		screen.DrawImage(img, opts)

		img.Dispose()
	})

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		img := sources.UseSources().Images().GetCopyOfImage("system/buttons/button")
		m := sources.UseSources().Metadata().GetMetadata("system/buttons/button_join_game")
		ms := m.GetMargins()
		s := m.GetScale()

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(ms.X, ms.Y)
		opts.GeoM.Scale(s.X, s.Y)

		p := positioning.UsePositioning().Button()
		p.Init(img, m, f, m.Text.Symbols)
		p.Draw()

		screen.DrawImage(img, opts)
	})
}
