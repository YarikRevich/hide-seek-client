package positioning

import (
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
	"github.com/YarikRevich/HideSeek-Client/internal/core/text/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

type Base struct {
	index, indent int

	examined []string

	font     font.Face
	img      *ebiten.Image
	metadata *sources.Model
}

func (b *Base) updateMetrics() {
	b.index++
	if b.index > 0 {
		b.indent = b.font.Metrics().Ascent.Round() + 2
	}
}

func (b *Base) getPosition() (int, int) {
	b.updateMetrics()

	switch b.metadata.Text.Position {
	case sources.Center:
		return b.getCenterCoords()
	case sources.Left:
		return b.getLeftCoords()
	case sources.Right:
		return b.getRightCoords()
	}
	return 0, 0
}

func (b *Base) getSymbolSize() int {
	symbolSize, ok := b.font.GlyphAdvance(rune(b.examined[b.index][0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}
	return symbolSize.Round()
}

func (b *Base) getText() string {
	return b.examined[b.index]
}

func (b *Base) getCenterCoords() (int, int) {
	symbolSize := b.getSymbolSize()
	return (int(b.metadata.Size.Width) - symbolSize*len(b.examined[b.index])) / 2, int(b.metadata.Size.Height/2) + b.indent
}

func (b *Base) getLeftCoords() (int, int) {
	return 0, int(b.metadata.Size.Height/2) + b.indent
}

func (b *Base) getRightCoords() (int, int) {
	symbolSize := b.getSymbolSize()
	return (int(b.metadata.Size.Width) - symbolSize*len(b.examined[b.index])) / 2, 0
}

func (b *Base) next() bool {
	return b.index < len(b.examined)-1
}

func (b *Base) Init(i *ebiten.Image, m *sources.Model, f font.Face, t string) {
	if len(t) == 0 {
		return
	}

	b.examined = strings.Split(t, "\n")
	b.img = i
	b.index = -1
	b.font = f
	b.metadata = m

	if len(b.examined) != 1 {
		b.indent = -5
	}else{
		b.indent = 0
	}
}

func (b *Base) Draw() {
	for b.next() {
		x, y := b.getPosition()

		fc := color.NewColor().GetColor(b.metadata.Fonts.FontColor)
		text.Draw(b.img, b.getText(), b.font, x, y, fc)
	}
}
