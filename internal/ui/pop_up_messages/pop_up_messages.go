package popupmessages

import (
	"image/color"

	"github.com/YarikRevich/HideSeek-Client/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)
var (
	MessageChan = make(chan error, 20 * 1024)
)

func Draw() {
	select {
	case msg := <- MessageChan:
		render.SetTextToRender(func(screen *ebiten.Image) {
			text.Draw(screen, msg.Error(), basicfont.Face7x13, 40, 40, color.RGBA{0xff, 0x00, 0x00, 0xff})
		})
	default:
		return
	}
}
