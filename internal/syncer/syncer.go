package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/hajimehoshi/ebiten/v2"
)

func SyncConfValues(screen *ebiten.Image) {
	sw, sh := screen.Size()
	swf, shf := float64(sw), float64(sh)
	pw, ph := history.GetScreenSize()
	pwf, phf := float64(pw), float64(ph)

	if pw != 0 && ph != 0 {
		SyncMetadata(swf, shf, pwf, phf)
		SyncPC(swf, shf, pwf, phf)
	}
}
