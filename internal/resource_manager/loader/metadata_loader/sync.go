package metadataloader

import (
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/hajimehoshi/ebiten/v2"
)

func SyncMetadata(screen *ebiten.Image) {
	sw, sh := screen.Size()
	swf, shf := float64(sw), float64(sh)
	pw, ph := history.GetScreenSize()
	pwf, phf := float64(pw), float64(ph)

	if pw != 0 && ph != 0 {
		for _, v := range MetadataCollection {
			v.Margins.TopMargin = (v.Margins.TopMargin * shf) / phf
			v.Margins.LeftMargin = (v.Margins.LeftMargin * swf) / pwf

			v.Scale.CoefficiantY = (v.Scale.CoefficiantY * shf) / phf
			v.Scale.CoefficiantX = (v.Scale.CoefficiantX * swf) / pwf
		}
	}
}
