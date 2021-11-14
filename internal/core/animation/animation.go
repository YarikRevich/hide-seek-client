package animation

import "github.com/YarikRevich/HideSeek-Client/internal/core/objects"

func Animate(o *objects.Object) {
	m := o.GetMetadata().Origin

	o.Animation.FrameDelayCounter++
	o.Animation.FrameDelayCounter %= uint64(m.Animation.FrameDelay)
	if o.Animation.FrameDelayCounter == 0 {
		o.Animation.FrameCount++
		o.Animation.FrameCount %= uint64(m.Animation.FrameNum)
	}
}
