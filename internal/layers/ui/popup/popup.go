package popup

import (
	"image/color"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/render"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var resourceLoadingCheck *time.Ticker = time.NewTicker(time.Second * 2)

func Draw() {
	select {
	case <-resourceLoadingCheck.C:
		if !sources.UseSources().IsLoadingEnded() {
			notifications.PopUp.WriteInfoWithPopUpTime("Resources are loading!", 5)
		}
	default:
	}

	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		var offset int
		for _, m := range notifications.PopUp.Read() {
			text.Draw(screen, m.Message, basicfont.Face7x13, 40, 40+offset, color.RGBA{0xff, 0x00, 0x00, 0xff})
			offset += 40
		}
	})
}
