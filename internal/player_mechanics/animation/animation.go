package animation

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func WithAnimation(src *ebiten.Image, m *models.Animation, p *pc.Animation) *ebiten.Image {
	
	p.FrameDelayCounter++
	p.FrameDelayCounter %= p.FrameDelay
	if p.FrameDelayCounter == 0 {
		p.FrameCount++
		p.FrameCount %= uint32(m.FrameNum)
	}
	sx, sy := int((m.FrameX+float64(p.FrameCount))*m.FrameWidth), int(m.FrameY)

	return src.SubImage(image.Rect(sx, sy, sx+int(m.FrameWidth), sy+int(m.FrameHeight))).(*ebiten.Image)
}
