package resume

import (
	"image/color"
	"math/rand"

	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	scale    = 64
	starsNum = 1024
)

var fromx, fromy, tox, toy, brightness float64

func createStar() {
	s := screen.UseScreen()
	fromx = tox
	fromy = toy
	tox = rand.Float64() * scale * s.GetWidth()
	toy = rand.Float64() * scale * s.GetHeight()
	brightness = rand.Float64() * 0xff
}

func Draw() {
	render.UseRender().SetToRender(func(i *ebiten.Image) {
		s := screen.UseScreen()

		fromx = tox
		fromy = toy
		tox += (tox - x) / 32
		toy += (toy - y) / 32
		brightness += 1
		if 0xff < brightness {
			brightness = 0xff
		}
		if fromx < 0 || s.GetWidth()*scale < fromx || fromy < 0 || s.GetHeight()*scale < fromy {
			createStar()
		}

		color := color.RGBA{uint8(0xbb * brightness / 0xff),
			uint8(0xdd * brightness / 0xff),
			uint8(0xff * brightness / 0xff),
			0xff}
		ebitenutil.DrawLine(i, fromx/scale, fromy/scale, tox/scale, toy/scale, color)
	})
}
