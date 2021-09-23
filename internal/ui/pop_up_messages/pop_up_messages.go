package popupmessages

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
)
var (
	MessageChan = make(chan error, 20 * 1024)
)

func Draw(screen *ebiten.Image) {
	select {
	case msg := <- MessageChan:
		text.Draw(screen, msg.Error(), basicfont.Face7x13, 40, 40, color.RGBA{0xff, 0x00, 0x00, 0xff})
	default:
		return
	}
}
