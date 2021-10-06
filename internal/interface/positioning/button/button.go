package button

import (
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

func ChooseButtonTextPosition(f font.Face, text string, m models.Metadata)(int, int){
	if len(text) == 0{
		return 0, 0
	}
	
	switch m.Button.TextPosition{
	case models.Center:
		return getCenterCoords(f, text, m)
	case models.Left:
		return getLeftCoords(f, m)
	case models.Right:
		return getRightCoords(f, text, m)
	}
	return 0, 0
}

func getCenterCoords(f font.Face, text string, m models.Metadata) (int, int) {	
	a, ok := f.GlyphAdvance(rune(text[0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(m.Size.Width) - a.Round()*len(text)) / 2, int(m.Size.Height / 2)
}

func getRightCoords(f font.Face, text string, m models.Metadata) (int, int) {
	a, ok := f.GlyphAdvance(rune(text[0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(m.Size.Width) - a.Round()*len(text)), int(m.Size.Height / 2)
}

func getLeftCoords(f font.Face, m models.Metadata) (int, int) {
	return 0, int(m.Size.Height / 2)
}
