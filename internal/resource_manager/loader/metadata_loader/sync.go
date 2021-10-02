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

			t := v.Margins.TopMargin
			l := v.Margins.LeftMargin

			v.Margins.TopMargin = (t * shf) / phf
			v.Margins.LeftMargin = (l * swf) / pwf

			sy := v.Scale.CoefficiantY
			sx := v.Scale.CoefficiantX

			v.Scale.CoefficiantY = (sy * shf) / phf
			v.Scale.CoefficiantX = (sx * swf) / pwf
		}
	}
}
