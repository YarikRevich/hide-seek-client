package popupmessages

import (
	"image/color"
	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func Draw() {
	render.UseRender().SetToRender(func(screen *ebiten.Image) {
		for _, m := range notifications.PopUp.Read() {
			text.Draw(screen, m.Message, basicfont.Face7x13, 40, 40, color.RGBA{0xff, 0x00, 0x00, 0xff})
		}
	})
}
