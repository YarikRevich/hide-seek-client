package animation

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
)

func Animate(o *objects.Base) {
	o.Animation.FrameDelayCounter++
	fmt.Println(o.ModelCombination.Modified.Animation, "PUK")
	o.Animation.FrameDelayCounter %= uint64(o.ModelCombination.Modified.Animation.FrameDelay)
	if o.Animation.FrameDelayCounter == 0 {
		o.Animation.FrameCount++
		o.Animation.FrameCount %= uint64(o.ModelCombination.Modified.Animation.FrameNum)
	}
}
