package animation

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
)

func Animate(o *objects.Base) {
	o.Animation.FrameDelayCounter++
	o.Animation.FrameDelayCounter %= uint64(o.ModelCombination.Modified.Animation.FrameDelay)
	if o.Animation.FrameDelayCounter == 0 {
		o.Animation.FrameCount++
		o.Animation.FrameCount %= uint64(o.ModelCombination.Modified.Animation.FrameNum)
	}
}
