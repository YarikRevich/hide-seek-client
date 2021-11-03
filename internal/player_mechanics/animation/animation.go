package animation

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	"github.com/hajimehoshi/ebiten/v2"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func WithAnimation(o *objects.Object) *ebiten.Image {
	m := metadatacollection.GetMetadata(o.Path)

	o.Animation.FrameDelayCounter++
	o.Animation.FrameDelayCounter %= uint32(m.Animation.FrameDelay)
	if o.Animation.FrameDelayCounter == 0 {
		o.Animation.FrameCount++
		o.Animation.FrameCount %= uint32(m.Animation.FrameNum)
	}
	sx, sy := int((m.Animation.FrameX+float64(o.Animation.FrameCount))*m.Animation.FrameWidth), int(m.Animation.FrameY)
	return imagecollection.GetImage(o.Path).SubImage(image.Rect(sx, sy, sx+int(m.Animation.FrameWidth), sy+int(m.Animation.FrameHeight))).(*ebiten.Image)
}
