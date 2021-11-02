package text

import (
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type PositioningType int

const (
	Button PositioningType = iota
	Input
)

type positionSession struct {
	indent               int
	examined             []string
	actualExaminedLength int
	index                int
	font                 font.Face
	stickWidth           float64
	stickHeight          float64
	position             models.TextPosition
	posType              PositioningType
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

func (p *positionSession) getFontMetrics() fixed.Int26_6 {
	a, ok := p.font.GlyphAdvance(rune(p.examined[p.index][0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}
	return a
}

func (p *positionSession) GetText() string {
	if p.posType == Input {
		a := p.getFontMetrics()
		curr := a.Round()*len(p.examined[p.index]) + (a.Round() * 5)
		if p.stickWidth < float64(curr) {
			t := p.examined[p.index]
			return string(t[:len(t)-int((float64(curr)-p.stickWidth)/float64(a.Round()))] + "...")
		}
	}
	return p.examined[p.index]
}

func (p *positionSession) getCenterCoords() (int, int) {
	a := p.getFontMetrics()
	if p.posType == Input {
		curr := a.Round()*len(p.examined[p.index]) + (a.Round() * 5)
		if p.stickWidth < float64(curr) {
			t := p.examined[p.index]
			s := string(t[:len(t)-int((float64(curr)-p.stickWidth)/float64(a.Round()))] + "...")
			return (int(p.stickWidth) - a.Round()*len(s)) / 2, int(p.stickHeight/2) + p.indent
		}
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
	posType PositioningType, font font.Face, text string, sw, sh float64, pos models.TextPosition) IPositionSession {
	p := new(positionSession)

	if len(text) != 0 {
		p.examined = strings.Split(text, "\n")
		p.index = -1
		p.font = font
		p.stickWidth = sw
		p.stickHeight = sh
		p.position = pos
		p.posType = posType

		if len(p.examined) != 1 {
			p.indent = -5
		}
	}

	return p
}
