package lines

import (
	"image/color"
	"math/rand"

	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"golang.org/x/image/colornames"
)

type Line struct {
	x0, y0, x1, y1 float64
	brightness     float64
	color          color.RGBA
}

func (l *Line) Shift(x, y float64, scale float64) {
	s := screen.UseScreen()
	size := s.GetSize()

	l.x0 = l.x1
	l.y0 = l.y1
	l.x1 += (l.x1 - x) / 32
	l.y1 += (l.y1 - y) / 32
	l.brightness += 1
	if l.brightness > 0xff {
		l.brightness = 0xff
	}
	if l.x0 < 0 || size.X*scale < l.x0 || l.y0 < 0 || size.Y*scale < l.y0 {
		l.SetRandomValues(scale)
	}

}

func (l *Line) Raw() (float64, float64, float64, float64, color.RGBA) {
	return l.x0, l.y0, l.x1, l.y1,
		color.RGBA{
			l.color.R / uint8(l.brightness),
			l.color.G / uint8(l.brightness),
			l.color.B / uint8(l.brightness),
			l.color.A / uint8(l.brightness)}
}

//TO DO
// func NewLine(x0, y0, x1, y1 float64, color color.RGBA) *Line {
// 	return &Line{x0, y0, x1, y1, rand.Float64() * 0xff, color}
// }

func (l *Line) SetRandomValues(scale float64) {
	s := screen.UseScreen()
	size := s.GetSize()
	l.x1 = rand.Float64() * size.X * scale
	l.x0 = l.x1
	l.y1 = rand.Float64() * size.Y * scale
	l.y0 = l.y1
	l.brightness = rand.Float64() * 0xff
	l.color = colornames.White
}

func NewEmptyLine() *Line {
	return new(Line)
}

type LinesPool struct {
	pool []Line
}

func (lp *LinesPool) ForEach(c func(Line)) {
	for _, v := range lp.pool {
		c(v)
	}
}

func NewLinesPool(linesNum int) *LinesPool {
	return &LinesPool{pool: make([]Line, linesNum)}
}
