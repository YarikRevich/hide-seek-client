package positioning

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/text/color"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Input struct {
	Base
}

func (i *Input) getPosition() (int, int) {
	i.updateMetrics()

	switch i.metadata.Text.Position {
	case sources.Center:
		return i.getCenterCoords()
	case sources.Left:
		return i.getLeftCoords()
	case sources.Right:
		return i.getRightCoords()
	}
	return 0, 0
}

func (i *Input) getRowWidth(symbolSize int) float64 {
	return float64(symbolSize*len(i.examined[i.index]) + (symbolSize * 5))
}

func (i *Input) getPlaceholder(rowWidth float64, symbolSize int) string {
	t := i.examined[i.index]

	return string(t[:len(t)-int((rowWidth-i.metadata.GetSize().X)/float64(symbolSize))] + "...")
}

func (i *Input) getText() string {
	symbolSize := i.getSymbolSize()
	rowWidth := i.getRowWidth(symbolSize)
	if i.metadata.GetSize().X < rowWidth {
		return i.getPlaceholder(rowWidth, symbolSize)
	}
	return i.Base.getText()
}

func (i *Input) getCenterCoords() (int, int) {
	symbolSize := i.getSymbolSize()
	rowWidth := i.getRowWidth(symbolSize)
	s := i.metadata.GetSize()
	if s.X < rowWidth {
		q := i.getPlaceholder(rowWidth, symbolSize)
		return (int(s.X) - symbolSize*len(q)) / 2, int(s.Y/2) + i.indent
	}
	return i.Base.getCenterCoords()
}

func (i *Input) Draw() {
	for i.next() {
		x, y := i.getPosition()
		fc := color.NewColor().GetColor(i.metadata.Fonts.FontColor)
		text.Draw(i.img, i.getText(), i.font, x, y, fc)
	}
}

func NewInput() *Input {
	return new(Input)
}
