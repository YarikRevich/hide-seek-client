package animation

import "github.com/YarikRevich/HideSeek-Client/internal/core/objects"

type PCAnimation struct {}

func (pa *PCAnimation) Animate() {
	p := objects.UseObjects().PC()
	m := p.GetMetadata().Origin

	p.Animation.FrameDelayCounter++
	p.Animation.FrameDelayCounter %= uint32(m.Animation.FrameDelay)
	if p.Animation.FrameDelayCounter == 0 {
		p.Animation.FrameCount++
		p.Animation.FrameCount %= uint32(m.Animation.FrameNum)
	}
}

func NewPCAnimation() *PCAnimation {
	return new(PCAnimation)
}
