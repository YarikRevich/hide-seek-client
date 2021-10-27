package animation

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func WithAnimation(o *objects.Object) *ebiten.Image {
	o.Animation.FrameDelayCounter++
	o.Animation.FrameDelayCounter %= uint32(o.Metadata.Animation.FrameDelay)
	if o.Animation.FrameDelayCounter == 0 {
		o.Animation.FrameCount++
		o.Animation.FrameCount %= uint32(o.Metadata.Animation.FrameNum)
	}
	
	sx, sy := int((o.Metadata.Animation.FrameX+float64(o.Animation.FrameCount))*o.Metadata.Animation.FrameWidth), int(o.Metadata.Animation.FrameY)

	return o.Image.SubImage(image.Rect(sx, sy, sx+int(o.Metadata.Animation.FrameWidth), sy+int(o.Metadata.Animation.FrameHeight))).(*ebiten.Image)
}
