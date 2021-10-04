package animation

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func WithAnimation(src *ebiten.Image, m models.Metadata, p *pc.Animation) *ebiten.Image {
	p.FrameDelayCounter++
	p.FrameDelayCounter %= p.FrameDelay
	if p.FrameDelayCounter == 0 {
		p.FrameCount++
		p.FrameCount %= uint32(m.Animation.FrameNum)
	}
	sx, sy := int((m.Animation.FrameX+float64(p.FrameCount))*m.Animation.FrameWidth), int(m.Animation.FrameY)

	return src.SubImage(image.Rect(sx, sy, sx+int(m.Animation.FrameWidth), sy+int(m.Animation.FrameHeight))).(*ebiten.Image)
}
