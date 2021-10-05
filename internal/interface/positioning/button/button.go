package button

import (
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

func ChooseButtonTextPosition(f font.Face, m models.Metadata)(int, int){
	switch m.Button.TextPosition{
	case models.Center:
		return getCenterCoords(f, m)
	case models.Left:
		return getLeftCoords(f, m)
	case models.Right:
		return getRightCoords(f, m)
	}
	return 0, 0
}

func getCenterCoords(f font.Face, m models.Metadata) (int, int) {
	a, ok := f.GlyphAdvance(rune(m.Button.Text[0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(m.Size.Width) - a.Round()*len(m.Button.Text)) / 2, int(m.Size.Height / 2)
}

func getRightCoords(f font.Face, m models.Metadata) (int, int) {
	a, ok := f.GlyphAdvance(rune(m.Button.Text[0]))
	if !ok {
		logrus.Fatal("error happened getting metrics of font")
	}

	return (int(m.Size.Width) - a.Round()*len(m.Button.Text)), int(m.Size.Height / 2)
}

func getLeftCoords(f font.Face, m models.Metadata) (int, int) {
	return 0, int(m.Size.Height / 2)
}
