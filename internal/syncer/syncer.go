package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/hajimehoshi/ebiten/v2"
	"sync"
)

func SyncConfValues(screen *ebiten.Image) {
	sw, sh := screen.Size()
	swf, shf := float64(sw), float64(sh)
	pw, ph := history.GetScreenSize()
	pwf, phf := float64(pw), float64(ph)

	if pw != 0 && ph != 0 {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			SyncMetadata(swf, shf, pwf, phf)
			wg.Done()
		}()

		go func() {
			SyncPC(swf, shf, pwf, phf)
			wg.Done()
		}()

		wg.Wait()
	}

}
