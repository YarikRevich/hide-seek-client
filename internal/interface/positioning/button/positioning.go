package positioning

import (
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

type positionSession struct {
	indent      int
	examined    []string
	index       int
	font        font.Face
	stickWidth  float64
	stickHeight float64
	position    models.TextPosition
}

type IPositionSession interface {
	Next() bool
	GetPosition() (int, int)
	GetText() string
}

func (p *positionSession) Next() bool {
	return p.index < len(p.examined)-1
}

func (p *positionSession) GetPosition() (int, int) {
	p.index++
	if p.index > 0 {
		p.indent = p.font.Metrics().Ascent.Round() + 2
	}

	switch p.position {
	case models.Center:
		return p.getCenterCoords()
	case models.Left:
		return p.getLeftCoords()
	case models.Right:
		return p.getRightCoords()
	}
	return 0, 0
}

func (p *positionSession) GetText() string {
	return p.examined[p.index]
}

func (p *positionSession) getCenterCoords() (int, int) {
	a, ok := p.font.GlyphAdvance(rune(p.examined[p.index][0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(p.stickWidth) - a.Round()*len(p.examined[p.index])) / 2, int(p.stickHeight/2) + p.indent
}

func (p *positionSession) getRightCoords() (int, int) {
	a, ok := p.font.GlyphAdvance(rune(p.examined[p.index][0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(p.stickWidth) - a.Round()*len(p.examined[p.index])), int(p.stickHeight/2) + p.indent
}

func (p *positionSession) getLeftCoords() (int, int) {
	return 0, int(p.stickHeight/2) + p.indent
}

func NewPositionSession(
	font font.Face, text string, sw, sh, cw, ch float64, pos models.TextPosition) IPositionSession {
	p := new(positionSession)

	if len(text) != 0 {
		p.examined = strings.Split(text, "\n")
		p.index = -1
		p.font = font
		p.stickWidth = sw
		p.stickHeight = sh
		p.position = pos

		if len(p.examined) == 1{
			p.indent = (int(p.stickHeight/4) - font.Metrics().Ascent.Round())
		}else{
			p.indent = -5
		}
	}

	return p
}
