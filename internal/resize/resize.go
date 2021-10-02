package resize

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/history"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

func SyncCoordinates(screen *ebiten.Image) {
	sw, sh := screen.Size()
	pw, ph := history.GetScreenSize()
	fmt.Println(sw, sh, pw, ph)
	if pw != 0 && ph != 0 {
		for _, v := range metadataloader.MetadataCollection {

			t := int(v.Margins.TopMargin)
			l := int(v.Margins.LeftMargin)

			v.Margins.TopMargin = float64((t * sh) / ph)
			v.Margins.LeftMargin = float64((l * sw) / pw)
		}
	}
}

func SyncScale(screen *ebiten.Image){

}
